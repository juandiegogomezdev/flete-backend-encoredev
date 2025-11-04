package appbusiness

import (
	"context"
	"fmt"

	"encore.app/appservice/sharedapp"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

func (b *BusinessApp) GetAllUserMemberships(ctx context.Context, userID uuid.UUID) ([]sharedapp.Membership, error) {
	memberships, err := b.store.GetUserMemberships(ctx, userID)
	if err != nil {
		fmt.Println("Error getting user memberships:", err)
		return nil, &errs.Error{
			Code:    errs.Internal,
			Message: "failed to get user memberships",
		}
	}

	return memberships, nil
}
