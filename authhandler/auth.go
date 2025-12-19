package authhandler

import (
	"context"
	"fmt"

	"encore.app/authhandler/sharedauth"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

var secrets struct {
	ClerkSecretKey string
}

//encore:service
type AuthHandler struct {
	jwksClient *jwks.Client
}

func initAuthHandler() (*AuthHandler, error) {
	if secrets.ClerkSecretKey == "" {
		return nil, fmt.Errorf("CLERK_SECRET_KEY no está configurada")
	}

	clerk.SetKey(secrets.ClerkSecretKey)

	config := &clerk.ClientConfig{}
	config.Key = clerk.String(secrets.ClerkSecretKey)

	jwksClient := jwks.NewClient(config)
	return &AuthHandler{jwksClient: jwksClient}, nil
}

//encore:authhandler
func (s *AuthHandler) AuthHandler(ctx context.Context, p *Params) (auth.UID, *AuthData, error) {

	// %+v prints with field names: {Name:Arun Age:25}
	userSub, err := verifySessionToken(ctx, p.Token, s.jwksClient)

	fmt.Println("user: ", userSub, "err: ", err)

	if err != nil {
		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al autenticar el usuario.",
			Details: sharedauth.ErrorDetailsToken{TokenStatus: "not_implemented"},
		}

	}

	// Este error es controlado, ignorar esto. No se ha implementado toda la logica por eso se lanza error
	return "", nil, &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "Error al autenticar el usuario. ",
		Details: sharedauth.ErrorDetailsToken{TokenStatus: "not_implemented"},
	}
}

type Params struct {
	Token string `header:"authorization"`
	// Extract the authorization header
	MembershipIdSelected string `header:"X-Membership-ID"`
}

type AuthData struct {
	UserID       uuid.UUID
	SessionID    uuid.UUID
	MembershipID uuid.UUID
}

func verifySessionToken(ctx context.Context, sessionToken string, jwksClient *jwks.Client) (string, error) {
	unsafeClaims, err := jwt.Decode(ctx, &jwt.DecodeParams{
		Token: sessionToken,
	})
	if err != nil {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al decodificar el token de sesión.",
		}
	}

	jwk, err := jwt.GetJSONWebKey(ctx, &jwt.GetJSONWebKeyParams{
		KeyID:      unsafeClaims.KeyID,
		JWKSClient: jwksClient,
	})
	if err != nil {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "1- Error al obtener la clave pública para verificar el token de sesión.",
		}
	}

	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: sessionToken,
		JWK:   jwk,
	})

	if err != nil {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "2-Error al obtener la clave pública para verificar el token de sesión.",
		}
	}

	usr, err := user.Get(ctx, claims.Subject)
	if err != nil {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "Error al obtener el usuario desde Clerk.",
		}
	}

	fmt.Println("user first name: ", *usr.FirstName)
	fmt.Println("user last name: ", usr.LastName)

	fmt.Println("user email: ", usr.EmailAddresses[0].EmailAddress)

	return usr.ID, nil

}
