package employee

type Service interface {
	GetEmployeesWithFilter(page, limit int, name string) ([]*Employee, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetEmployeesWithFilter(page, limit int, name string) ([]*Employee, error) {
	return s.repo.FindWithFilter(page, limit, name)
}
