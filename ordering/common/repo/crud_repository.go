package repo

import (
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type CrudRepository[T any] struct {
	DB *gorm.DB
}

func (repo *CrudRepository[T]) Migrate() error {
	return repo.DB.AutoMigrate(new(T))
}
func (repo *CrudRepository[T]) FindByID(ID uint) (*T, error) {
	result := new(T)
	trans := repo.DB.First(&result, ID)
	return result, trans.Error
}
func (repo *CrudRepository[T]) FindOne(conditions ...interface{}) (*T, error) {
	result := new(T)
	trans := repo.DB.First(&result, conditions...)
	return result, trans.Error
}
func (repo *CrudRepository[T]) Find(conditions ...interface{}) ([]*T, error) {
	results := make([]*T, 0)
	trans := repo.DB.Find(&results, conditions...)
	if trans.Error != nil {
		return nil, trans.Error
	}
	return results, nil
}
func (repo *CrudRepository[T]) Insert(data *T) error {
	trans := repo.DB.Save(data)
	return trans.Error
}
func (repo *CrudRepository[T]) Edit(data *T) error {
	trans := repo.DB.Save(data)
	return trans.Error
}
func (repo *CrudRepository[T]) Delete(data *T) error {
	trans := repo.DB.Delete(data)
	return trans.Error
}
