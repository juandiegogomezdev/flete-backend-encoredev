package organizationbusiness

import (
	"context"

	"encore.app/organizationoservice/shared/types"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

func (b *OrganizationBusiness) CreateOrganization(ctx context.Context, name string, ownerId auth.UID) (types.Membership, error) {
	// Check if organization name is available for the user
	available, err := b.s.IsAvailableOrganizationName(ctx, name, ownerId)
	if err != nil {
		return types.Membership{}, &errs.Error{
			Message: "Error al comprobar la disponibilidad del nombre de la organización",
			Code:    errs.Internal,
		}
	}
	if !available {
		return types.Membership{}, &errs.Error{
			Message: "El nombre de la organización ya está en uso",
			Code:    errs.AlreadyExists,
		}
	}

	// Get the role ID for "Owner"
	roleID, err := b.s.GetRoleIdByName(ctx, "Owner")
	if err != nil {
		return types.Membership{}, &errs.Error{
			Message: "Error al obtener el rol de propietario",
			Code:    errs.Internal,
		}
	}

	// Create organization and membership
	membership, err := b.s.CreateOrganizationAndMembership(ctx, name, ownerId, roleID)
	if err != nil {
		return types.Membership{}, &errs.Error{
			Message: "Error al crear la organización y la membresía",
			Code:    errs.Internal,
		}
	}
	return membership, nil
}
