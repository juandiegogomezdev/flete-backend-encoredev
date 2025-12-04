package authService

import (
	"context"
	"net/http"
	"strings"

	"encore.app/authService/sharedauth"
	"encore.app/pkg/myjwt"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

var secrets struct {
	JWT_SECRET_KEY string
}

var tokenizer = myjwt.NewJWTTokenizer(secrets.JWT_SECRET_KEY)

func parseMembershipToken(token string) (*myjwt.MembershipClaims, error) {
	claims, status := tokenizer.ParseMembershipToken(token)
	switch status {
	case myjwt.TokenStatusExpired:
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "El token ha expirado",
			Details: sharedauth.ErrorDetailsToken{TokenStatus: string(myjwt.TokenStatusExpired)},
		}
	case myjwt.TokenStatusInvalid:
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "El token no es válido",
			Details: sharedauth.ErrorDetailsToken{TokenStatus: string(myjwt.TokenStatusInvalid)},
		}
	case myjwt.TokenStatusValid:
		return claims, nil
	}
	return nil, &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "Error al validar el token",
		Details: sharedauth.ErrorDetailsToken{TokenStatus: string(myjwt.TokenStatusInvalid)},
	}
}

type MyAuthParams struct {
	// Extract the auth token from either the cookie or the Authorization header
	SessionCookie *http.Cookie `cookie:"auth_token"`
	// Extract the authorization header
	AuthorizationHeader string `header:"Authorization"`
}

//encore:authhandler
func AuthHandler(ctx context.Context, p *MyAuthParams) (auth.UID, *AuthData, error) {
	// Extract the token from the request
	token, err := extractToken(p)
	if err != nil {
		return "", nil, err
	}

	// Parse the token and get the claims
	claims, err := parseMembershipToken(token)
	if err != nil {
		return "", nil, err
	}

	authData := &AuthData{
		UserID:       claims.UserID,
		SessionID:    claims.SessionID,
		MembershipID: claims.MembershipID,
	}
	return auth.UID(claims.UserID.String()), authData, nil
}

type AuthData struct {
	UserID       uuid.UUID
	SessionID    uuid.UUID
	MembershipID uuid.UUID
}

// Extract token from AuthParams
func extractToken(p *MyAuthParams) (token string, err error) {
	// Verify if the token is in the cookie
	if p.SessionCookie != nil {
		if p.SessionCookie.Value == "" {
			return "", &errs.Error{
				Code:    errs.Unauthenticated,
				Message: "cookie de sesion vacia",
			}
		}
		return p.SessionCookie.Value, nil
	}

	// Verify if the token is in the Authorization header
	if p.AuthorizationHeader != "" {
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

	return "", &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "se requiere autenticacion: proporciona una cookie o un header de autorizacion",
		Details: sharedauth.ErrorDetailsToken{TokenStatus: string(myjwt.TokenStatusInvalid)},
	}
}
