package appService

// Send the memberships of an user
// encore:api auth method=GET path=/memberships
// func (s *ServiceApp) GetUserMemberships(ctx context.Context) (GetUserMembershipsResponse, error) {
// 	memberships, err := s.b.GetAllUserMemberships(ctx, auth.Data().(*authhandler.AuthData).UserID)
// 	if err != nil {
// 		return GetUserMembershipsResponse{}, err
// 	}
// 	time.Sleep(3 * time.Second)
// 	return GetUserMembershipsResponse{Memberships: memberships}, nil
// }

// type GetUserMembershipsResponse struct {
// 	Memberships []sharedapp.Membership `json:"data"`
// }
