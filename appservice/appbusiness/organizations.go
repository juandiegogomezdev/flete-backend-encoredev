package appbusiness

import (
	"context"
	"fmt"
	"time"

	"encore.app/appservice/sharedapp"
	"encore.app/pkg/utils"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

// Get all memberships of a user
// func (b *BusinessApp) GetAllUserMemberships(ctx context.Context, userID uuid.UUID) ([]appstore.ResUserOrganizationStore, error) {

// 	organizations, err := b.store.GetAllUserMemberships(ctx, userID)
// 	if err != nil {
// 		return nil, &errs.Error{
// 			Code:    errs.Internal,
// 			Message: "Error al obtener las organizaciones del usuario",
// 		}
// 	}

// 	return organizations, nil
// }

// Create a company organization
func (b *BusinessApp) CreateCompanyOrganization(ctx context.Context, userID uuid.UUID, name string) (sharedapp.Membership, error) {
	// Get all organizations for this user
	existingOrgs, err := b.store.GetAllUserOrganizations(ctx, userID)
	if err != nil {
		fmt.Println("Error checking existing organizations:", err)
		return sharedapp.Membership{}, &errs.Error{
			Code:    errs.Internal,
			Message: "Error al crear la organización de la empresa",
		}
	}

	// Only is posible create 4 organizations per user and the user need to have at least the
	// personal organization created
	if len(existingOrgs) >= 4 && len(existingOrgs) != 0 {
		return sharedapp.Membership{}, &errs.Error{
			Code:    errs.FailedPrecondition,
			Message: "No es posible crear más organizaciones. Límite alcanzado.",
		}
	}

	// Generate the new organization ID and membership ID
	orgID, err := utils.GenerateUUID()
	if err != nil {
		return sharedapp.Membership{}, err
	}
	memID, err := utils.GenerateUUID()
	if err != nil {
		return sharedapp.Membership{}, err
	}

	// Create the company organization struct
	createOrganization := sharedapp.CreateOrganizationStruct{
		Name:    name,
		OrgID:   orgID,
		OwnerID: userID,
	}

	// Create the membership organization struct
	createMembership := sharedapp.CreateOwnerMembershipStruct{
		MemID:     memID,
		OrgID:     orgID,
		UserID:    userID,
		Status:    "active",
		CreatedBy: userID,
		RoleID:    uuid.Nil,
	}

	// Create the organization and the membership in a transaction
	err = b.store.CreateOrgAndMembership(ctx, createOrganization, createMembership)
	if err != nil {
		return sharedapp.Membership{}, &errs.Error{
			Code:    errs.Internal,
			Message: "Error al crear el trabajo como independiente",
		}
	}

	membership := sharedapp.Membership{
		ID:        memID,
		Status:    "active",
		CreatedAt: time.Now(),
		OrgName:   name,
		RoleName:  "owner",
	}
	return membership, nil
}
