package appService

import (
	"context"

	"encore.app/appservice/shared"
	"encore.dev/beta/auth"
)

//encore:api public method=GET path=/org/hello
func (s *ServiceApp) Hello(ctx context.Context) (*responseHello, error) {
	return &responseHello{Message: "Hello, World!"}, nil
}

//encore:api public method=GET path=/org
func (s *ServiceApp) GetAllOrganizations(ctx context.Context) (responseGetAllOrganizations, error) {
	return responseGetAllOrganizations{}, nil
}

type responseGetAllOrganizations struct {
}

type responseHello struct {
	Message string `json:"message"`
}

// Create a personal organization for the user
// encore:api auth method=POST path=/org/personal
func (s *ServiceApp) CreatePersonalOrg(ctx context.Context) (*CreateOrgResponse, error) {
	data := auth.Data().(*AuthData)
	membership, err := s.b.CreatePersonalOrganization(ctx, data.UserID, "Trabajo independiente")
	if err != nil {
		return nil, err
	}
	return &CreateOrgResponse{Memberships: membership}, nil
}

type CreateOrgResponse struct {
	Memberships shared.Membership `json:"data"`
}

// Create a company organization for the user
// encore:api auth method=POST path=/org/company
func (s *ServiceApp) CreateCompanyOrg(ctx context.Context, req *reqCreateCompanyOrgRequest) (*CreateOrgResponse, error) {
	data := auth.Data().(*AuthData)
	membership, err := s.b.CreateCompanyOrganization(ctx, data.UserID, req.CompanyName)
	if err != nil {
		return nil, err
	}
	return &CreateOrgResponse{Memberships: membership}, nil
}

type reqCreateCompanyOrgRequest struct {
	CompanyName string `json:"companyName"`
}
