package organizationoservice

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"encore.app/authhandler"
	"encore.app/databaseservice"
	"encore.app/organizationoservice/organizationbusiness"
	"encore.app/organizationoservice/organizationstore"
	"encore.app/organizationoservice/shared/types"
	"encore.app/pkg/utils"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var primaryDB = sqldb.Named("primary_db")

var primaryDBX = sqlx.NewDb(primaryDB.Stdlib(), "postgres")

//encore:service
type OrganizationService struct {
	b *organizationbusiness.OrganizationBusiness
}

func initOrganizationService() (*OrganizationService, error) {
	store := organizationstore.NewOrganizationStore(primaryDB, primaryDBX)
	business := organizationbusiness.NewOrganizationBusiness(store)
	return &OrganizationService{b: business}, nil
}

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
func (s *OrganizationService) UploadLogoOrganization(w http.ResponseWriter, r *http.Request) {

	userId, _ := auth.UserID()
	userData := auth.Data().(*authhandler.AuthData)

	// Get the organization id
	orgId, err := s.b.GetOrganizationByMembershipId(r.Context(), userData.MembershipID)

	// Limit the request to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 11<<20)

	// 1. Parse the multipart form
	err = utils.ParseMultipartForm(r)
	if err != nil {
		http.Error(w, "La imagen es demasiado pesada", http.StatusRequestEntityTooLarge)
		return
	}
	defer r.MultipartForm.RemoveAll()

	// Read the logo if exists
	file, header, err := utils.ExtractFileFromMultipartform(r, "logo", 10*utils.MB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	//
	uuidFile, err := utils.MustNewUUID()
	if err != nil {
		log.Println("Error generating uuidFile for logo name")
		http.Error(w, "Error al procesar la imagen")
		return
	}
	key := fmt.Sprintf("%s/%s/%s/logo/%s", userId, orgId, uuidFile)
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

	// Save the organization in the database
	membership, err := s.b.CreateOrganization(r.Context(), name, key, userId)
	if err != nil {
		errs.HTTPError(w, err)
		return
	}

	fmt.Println("Membership created:", membership)

	// Respond success

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Organización creada con éxito por el usuario " + string(userId)))

}
