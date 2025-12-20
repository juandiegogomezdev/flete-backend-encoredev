package authhandler

import (
	"context"
	"fmt"

	"encore.app/authhandler/authbusiness"
	"encore.app/authhandler/authstore"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
	"encore.dev/types/uuid"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

var primaryDB = sqldb.Named("primary_db")

var secrets struct {
	ClerkSecretKey string
}

//encore:service
type AuthHandler struct {
	jwksClient *jwks.Client
	b          *authbusiness.AuthBusiness
}

func initAuthHandler() (*AuthHandler, error) {

	// Clerk setup
	if secrets.ClerkSecretKey == "" {
		return nil, fmt.Errorf("CLERK_SECRET_KEY no está configurada")
	}

	clerk.SetKey(secrets.ClerkSecretKey)

	config := &clerk.ClientConfig{}
	config.Key = clerk.String(secrets.ClerkSecretKey)

	jwksClient := jwks.NewClient(config)

	// Initialize AuthBusiness with AuthStore
	store := authstore.NewAuthStore(primaryDB)
	business := authbusiness.NewAuthBusiness(store)

	return &AuthHandler{jwksClient: jwksClient, b: business}, nil
}

//encore:authhandler
func (s *AuthHandler) AuthHandler(ctx context.Context, p *Params) (auth.UID, *AuthData, error) {

	// Verify session token and get user info
	usr, err := verifySessionToken(ctx, p.Token, s.jwksClient)
	if err != nil {
		return "", nil, err
	}

	// Ensure user exists in our database
	err = s.b.CreateUserIfNotExists(ctx, usr)
	if err != nil {
		return "", nil, err
	}

	// Return auth data with selected membership ID
	membershipID := uuid.FromStringOrNil(p.MembershipIdSelected)
	return auth.UID(usr.ID), &AuthData{
		MembershipID: membershipID,
	}, nil

}

type Params struct {
	Token                string `header:"authorization"`
	MembershipIdSelected string `header:"X-Membership-ID"`
}

type AuthData struct {
	MembershipID uuid.UUID
}

func verifySessionToken(ctx context.Context, sessionToken string, jwksClient *jwks.Client) (*clerk.User, error) {
	unsafeClaims, err := jwt.Decode(ctx, &jwt.DecodeParams{
		Token: sessionToken,
	})
	if err != nil {
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al decodificar el token de sesión.",
		}
	}

	jwk, err := jwt.GetJSONWebKey(ctx, &jwt.GetJSONWebKeyParams{
		KeyID:      unsafeClaims.KeyID,
		JWKSClient: jwksClient,
	})
	if err != nil {
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "1- Error al obtener la clave pública para verificar el token de sesión.",
		}
	}

	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: sessionToken,
		JWK:   jwk,
	})

	if err != nil {
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "2-Error al obtener la clave pública para verificar el token de sesión.",
		}
	}

	usr, err := user.Get(ctx, claims.Subject)
	if err != nil {
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al obtener el usuario desde Clerk.",
		}
	}

	return usr, nil

}
