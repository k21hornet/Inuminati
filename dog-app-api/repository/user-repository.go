package repository

import (
	"dog-app-api/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// インターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	GetUserById(user *model.User, userId uint) error
	UpdateIcon(user *model.User, userId uint, s3URL string) error
}

// 構造体
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ、dbへ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// emailからユーザーを取得する
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	//dbから引数で受け取ったemailのuserが存在するか
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザー作成
func (ur *userRepository) CreateUser(user *model.User) error {
	// 引数でユーザーオブジェクトのポインタを取得し、ユーザーを作成(新規ユーザーで書き換え)
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザー取得
func (ur *userRepository) GetUserById(user *model.User, userId uint) error {
	if err := ur.db.Where("id=?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザーアイコン更新
func (ur *userRepository) UpdateIcon(user *model.User, userId uint, s3URL string) error {
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?", userId).Update("icon", s3URL)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
