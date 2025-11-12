package services

type RoleService struct {
}

func NewRoleService() (*RoleService, error) {
	return &RoleService{}, nil
}

func MustRoleService() *RoleService {
	s, e := NewRoleService()
	if e != nil {
		panic(e)
	}
	return s
}
