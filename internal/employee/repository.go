package employee

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindWithFilter(page, limit int, name string) ([]*Employee, error)
	CountWithFilter(name string) (int64, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) FindWithFilter(page, limit int, name string) ([]*Employee, error) {
	var employees []*Employee
	offset := (page - 1) * limit

	query := r.DB.Model(&Employee{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	err := query.Limit(limit).Offset(offset).Find(&employees).Error
	return employees, err
}

func (r *repository) CountWithFilter(name string) (int64, error) {
	var count int64
	query := r.DB.Model(&Employee{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	err := query.Count(&count).Error
	return count, err
}
