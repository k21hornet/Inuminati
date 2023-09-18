package repository

import (
	"dog-app-api/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// インターフェース
type IDogRepository interface {
	GetAllUsersDogs(dogs *[]model.Dog) error
	GetAllDogs(dogs *[]model.Dog, userId uint) error
	GetDogById(dog *model.Dog, dogId uint) error
	CreateDog(dog *model.Dog) error
	UpdateDog(dog *model.Dog, userId uint, dogId uint) error
	DeleteDog(userId uint, dogId uint) error
}

// 構造体
type dogRepository struct {
	db *gorm.DB
}

// コンストラクタ、dbへ
func NewDogRepository(db *gorm.DB) IDogRepository {
	return &dogRepository{db}
}

func (dr *dogRepository) GetAllUsersDogs(dogs *[]model.Dog) error {
	if err := dr.db.Joins("User").Order("created_at DESC").Find(dogs).Error; err != nil {
		return err
	}
	return nil
}

func (dr *dogRepository) GetAllDogs(dogs *[]model.Dog, userId uint) error {
	if err := dr.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Find(dogs).Error; err != nil {
		return err
	}
	return nil
}

func (dr *dogRepository) GetDogById(dog *model.Dog, dogId uint) error {
	if err := dr.db.Joins("User").First(dog, dogId).Error; err != nil {
		return err
	}
	return nil
}

func (dr *dogRepository) CreateDog(dog *model.Dog) error {
	if err := dr.db.Create(dog).Error; err != nil {
		return err
	}
	return nil
}

func (dr *dogRepository) UpdateDog(dog *model.Dog, userId uint, dogId uint) error {
	result := dr.db.Model(dog).Clauses(clause.Returning{}).Where("id=? AND user_id=?", dogId, userId).Update("img", dog.Img)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (dr *dogRepository) DeleteDog(userId uint, dogId uint) error {
	result := dr.db.Where("id=? AND user_id=?", dogId, userId).Delete(&model.Dog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
