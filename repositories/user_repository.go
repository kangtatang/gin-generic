package repositories

import (
	"gorm.io/gorm"
)

// GenericRepository adalah fungsi generik untuk operasi CRUD dengan GORM
type GenericRepository[T any] struct {
	DB *gorm.DB
}

// GetAll mengembalikan semua data dari tabel
func (r *GenericRepository[T]) GetAll() ([]T, error) {
	var results []T
	if err := r.DB.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// GetByID mencari data berdasarkan ID
func (r *GenericRepository[T]) GetByID(id uint) (*T, error) {
	var result T
	if err := r.DB.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// Create membuat data baru
func (r *GenericRepository[T]) Create(data *T) error {
	return r.DB.Create(data).Error
}
