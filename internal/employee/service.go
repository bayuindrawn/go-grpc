package employee

type Service interface {
	GetAllEmployees() ([]*Employee, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllEmployees() ([]*Employee, error) {
	return s.repo.FindAll()
}
