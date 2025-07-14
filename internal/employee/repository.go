package employee

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]*Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]*Employee, error) {
	var employees []*Employee
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
