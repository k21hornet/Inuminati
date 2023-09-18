package usecase

import (
	"dog-app-api/model"
	"dog-app-api/repository"
	"dog-app-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// インターフェース
type IUserUsecase interface {
	// ユーザーとエラーを返す
	SignUp(user model.User) (model.UserResponse, error)
	// jwtのstringを返す
	Login(user model.User) (string, error)
	//
	GetDogById(userId uint) (model.UserResponse, error)
	UpdateIcon(user model.User, userId uint, s3URL string) (model.UserResponse, error)
}

// 構造体
type userUsecase struct {
	// usecaseはrepositoryのインターフェースのみに依存
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// コンストラクタ、repositoryへ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

// 新規登録
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// validate
	if err := uu.uv.NewUserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	// bcryptのパッケージで平文のパスワードをハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{UserName: user.UserName, Email: user.Email, Password: string(hash), Icon: "https://dog-app-upload-dev-1.s3.amazonaws.com/icons/default-icon.png"}
	// repositoryへ委譲
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:       newUser.ID,
		UserName: newUser.UserName,
		Email:    newUser.Email,
		Icon:     newUser.Icon,
	}
	return resUser, nil
}

// ログイン
func (uu *userUsecase) Login(user model.User) (string, error) {
	// validate
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	//emailがデータベースに存在するか否か
	storedUser := model.User{} //空のユーザーオブジェクト
	// repositoryへ委譲
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// パスワードの検証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// jwt, payload:userid,有効期限(12h)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	// jwt生成
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ユーザーを探す
func (uu *userUsecase) GetDogById(userId uint) (model.UserResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserById(&user, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}
	return resUser, nil
}

func (uu *userUsecase) UpdateIcon(user model.User, userId uint, s3URL string) (model.UserResponse, error) {
	if err := uu.ur.UpdateIcon(&user, userId, s3URL); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}
	return resUser, nil
}
