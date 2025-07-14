package employee

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	FindWithFilter(ctx context.Context, page, limit int, name string) ([]*Employee, error)
	CountWithFilter(ctx context.Context, name string) (int64, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) FindWithFilter(ctx context.Context, page, limit int, name string) ([]*Employee, error) {
	var employees []*Employee
	offset := (page - 1) * limit

	query := r.DB.WithContext(ctx).Model(&Employee{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Limit(limit).Offset(offset).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *repository) CountWithFilter(ctx context.Context, name string) (int64, error) {
	var count int64
	query := r.DB.WithContext(ctx).Model(&Employee{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	err := query.Count(&count).Error
	return count, err
}
