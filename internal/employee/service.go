package employee

import "context"

type Service interface {
	GetEmployeesWithFilter(ctx context.Context, page, limit int, name string) ([]*Employee, int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetEmployeesWithFilter(ctx context.Context, page, limit int, name string) ([]*Employee, int64, error) {
	employees, err := s.repo.FindWithFilter(ctx, page, limit, name)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.repo.CountWithFilter(ctx, name)
	if err != nil {
		return nil, 0, err
	}

	return employees, count, nil
}
