package authhandler

import (
	"context"
	"fmt"
	"strings"

	"encore.app/authhandler/sharedauth"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
)

var secrets struct {
	ClerkSecretKey string
}

//encore:service
type AuthHandler struct{}

func initAuthHandler() (*AuthHandler, error) {
	clerk.SetKey(secrets.ClerkSecretKey)
	return &AuthHandler{}, nil
}

//encore:authhandler
func (s *AuthHandler) AuthHandler(ctx context.Context, p *MyAuthParams) (auth.UID, *AuthData, error) {
	// Extract the token from the request
	token, err := extractToken(p)
	if err != nil {
		return "", nil, err
	}

	fmt.Println("Token: ", token)

	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: token,
	})

	fmt.Println("claims: ", claims, "err: ", err)

	if err != nil {
		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al autenticar el usuario.",
			Details: sharedauth.ErrorDetailsToken{TokenStatus: "not_implemented"},
		}

	}

	return "", nil, &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "Error al autenticar el usuario. ",
		Details: sharedauth.ErrorDetailsToken{TokenStatus: "not_implemented"},
	}
}

type MyAuthParams struct {
	// Extract the authorization header
	AuthorizationHeader string `header:"Authorization"`
}

type AuthData struct {
	UserID       uuid.UUID
	SessionID    uuid.UUID
	MembershipID uuid.UUID
}

// Extract token from AuthParams
func extractToken(p *MyAuthParams) (token string, err error) {
	if p.AuthorizationHeader == "" {
		fmt.Errorf("No authentication token provided")
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al autenticar el usuario.",
		}
	}

	if after, found := strings.CutPrefix(p.AuthorizationHeader, "Bearer "); found {
		token = strings.TrimSpace(after)
		if token != "" {
			return token, nil
		}
	}
	return "", &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "Formato de header de autorizacion invalido",
		Details: sharedauth.ErrorDetailsToken{TokenStatus: "invalid_token"},
	}

}
