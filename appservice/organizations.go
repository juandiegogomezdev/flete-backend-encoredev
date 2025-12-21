package appService

import (
	"context"
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

// Create a company organization for the user
// encore:api auth method=POST path=/org/company
// func (s *ServiceApp) CreateCompany(ctx context.Context, req *reqCreateCompanyOrgRequest) (*CreateOrgResponse, error) {
// 	data := auth.Data().(*authhandler.AuthData)
// 	membership, err := s.b.CreateCompanyOrganization(ctx, data.UserID, req.CompanyName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &CreateOrgResponse{Memberships: membership}, nil
// }

// type reqCreateCompanyOrgRequest struct {
// 	CompanyName string `json:"companyName"`
// }

// type CreateOrgResponse struct {
// 	Memberships sharedapp.Membership `json:"data"`
// }
