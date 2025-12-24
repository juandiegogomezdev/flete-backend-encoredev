package organizationoservice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"encore.app/authhandler"
	"encore.app/databaseservice"
	types "encore.app/organizationoservice/shared/types"
	"encore.app/pkg/utils"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/objects"
)

//encore:api auth method=POST path=/create-organization
func (s *OrganizationService) CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (types.Membership, error) {
	userId, _ := auth.UserID()
	membership, err := s.b.CreateOrganization(ctx, req.Name, userId)
	if err != nil {
		return types.Membership{}, err
	}
	return membership, nil
}

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}
type CreateOrganizationResponse struct {
	Membership types.Membership `json:"membership"`
}

//encore:api auth raw method=POST path=/upload-logo
func (s *OrganizationService) UploadOrganizationLogo(w http.ResponseWriter, r *http.Request) {

	userId, _ := auth.UserID()
	userData := auth.Data().(*authhandler.AuthData)
	org, err := s.b.GetOrganizationByMembershipId(r.Context(), userData.MembershipID)

	// Delete previous logo in database
	err = s.b.DeleteOrganizationLogo(r.Context(), org.ID)
	if err != nil {
		errs.HTTPError(w, err)
		return
	}

	// Delete previous logo in storage
	err = databaseservice.PrimaryBucketPublic.Remove(r.Context(), *org.ImageKey)
	if err != nil && !errors.Is(err, objects.ErrObjectNotFound) {
		log.Println("Error deleting previous logo from storage:", err)
		errs.HTTPError(w, err)
		return
	}

	// Parse multipart form
	r.Body = http.MaxBytesReader(w, r.Body, 11<<20)
	err = utils.ParseMultipartForm(r)
	if err != nil {
		http.Error(w, "La imagen es demasiado pesada", http.StatusRequestEntityTooLarge)
		return
	}
	defer r.MultipartForm.RemoveAll()

	// Read the logo file
	file, header, err := utils.ExtractFileFromMultipartform(r, "logo", 10*utils.MB)
	if err != nil {
		errs.HTTPError(w, err)
		return
	}
	defer file.Close()

	// Generate key: userId/orgId/logo/uuidFile
	uuidFile, err := utils.MustNewUUID()
	if err != nil {
		errs.HTTPError(w, err)
		return
	}
	uuidFileStr := fmt.Sprintf("%s.%s", uuidFile.String(), utils.GetExtensionFromFilename(header.Filename))
	key := fmt.Sprintf("%s/%s/logo/%s", userId, org.ID, uuidFileStr)

	// Create writer and upload the file
	writer := databaseservice.PrimaryBucketPublic.Upload(r.Context(), key)
	_, err = io.Copy(writer, file)
	if err != nil {
		writer.Abort(err)
		errs.HTTPError(w, err)
		return
	}
	if err := writer.Close(); err != nil {
		errs.HTTPError(w, err)
		return
	}

	// Save the logo key in the organization record
	err = s.b.SaveOrganizationLogoKey(r.Context(), org.ID, key)
	if err != nil {
		// Delete the logo in storage
		err = databaseservice.PrimaryBucketPublic.Remove(r.Context(), *org.ImageKey)
		if err != nil && !errors.Is(err, objects.ErrObjectNotFound) {
			log.Println("Error deleting previous logo from storage:", err)
		}
		errs.HTTPError(w, err)
		return
	}

	// Respond success

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logo cargado satisfactoriamente " + string(userId)))

}
