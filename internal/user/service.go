package user

type Service interface {
	Login(entity *Login) (*User, error)
}

type service struct {
	repo Repository
}

func (s service) Login(entity *Login) (*User, error) {
	return s.repo.GetUser(entity)
}

func newService(repo Repository) Service {
	return &service{repo}
}
