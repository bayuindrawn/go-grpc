package employee

type Service interface {
	GetEmployeesWithFilter(page, limit int, name string) ([]*Employee, int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetEmployeesWithFilter(page, limit int, name string) ([]*Employee, int64, error) {
	employees, err := s.repo.FindWithFilter(page, limit, name)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.repo.CountWithFilter(name)
	if err != nil {
		return nil, 0, err
	}

	return employees, count, nil
}
