package authbusiness

import (
	"context"
	"fmt"

	"encore.app/authhandler/authstore"
	"encore.dev/beta/errs"
	"github.com/clerk/clerk-sdk-go/v2"
)

type AuthBusiness struct {
	store *authstore.AuthStore
}

func NewAuthBusiness(store *authstore.AuthStore) *AuthBusiness {
	return &AuthBusiness{store: store}
}

func (s *AuthBusiness) CreateUserIfNotExists(ctx context.Context, usr *clerk.User) error {
	exists, err := s.store.CheckUserExists(ctx, usr.ID)
	if err != nil {
		fmt.Println("Error checking if user exists: ", err)
		return &errs.Error{
			Code:    errs.Internal,
			Message: "Error al verificar si el usuario existe en la base de datos.",
		}
	}

	if exists {
		return nil
	}

	email := ""

	if usr.EmailAddresses != nil && len(usr.EmailAddresses) > 0 {
		email = usr.EmailAddresses[0].EmailAddress
	}

	phone := ""
	if usr.PhoneNumbers != nil && len(usr.PhoneNumbers) > 0 {
		phone = usr.PhoneNumbers[0].PhoneNumber
	}

	user := authstore.UserCreateParams{
		ID:        usr.ID,
		Email:     email,
		FirstName: *usr.FirstName,
		LastName:  *usr.LastName,
		Phone:     phone,
	}

	fmt.Println("new User", user)

	err = s.store.CreateUser(ctx, &user)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return &errs.Error{
			Code:    errs.Internal,
			Message: "Error al crear el usuario en la base de datos.",
		}
	}
	return nil
}
