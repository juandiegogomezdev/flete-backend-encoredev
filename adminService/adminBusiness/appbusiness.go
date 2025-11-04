package adminBusiness

import "encore.app/adminService/adminstore"

type AdminBusiness struct {
	s *adminstore.AdminStore
}

func NewAdminBusiness(s *adminstore.AdminStore) *AdminBusiness {
	return &AdminBusiness{s: s}
}
