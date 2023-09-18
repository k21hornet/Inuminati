package usecase

import (
	"dog-app-api/model"
	"dog-app-api/repository"
	"dog-app-api/validator"
)

// インターフェース
type IDogUsecase interface {
	GetAllUsersDogs() ([]model.DogResponse, error)
	GetAllDogs(userId uint) ([]model.DogResponse, error)
	GetDogById(dogId uint) (model.DogResponse, error)
	CreateDog(dog model.Dog) (model.DogResponse, error)
	UpdateDog(dog model.Dog, userId uint, dogId uint) (model.DogResponse, error)
	DeleteDog(userId uint, dogId uint) error
}

// 構造体
type dogUsecase struct {
	dr repository.IDogRepository
	dv validator.IDogValidator
}

// コンストラクタ、repositoryへ
func NewDogUsecase(dr repository.IDogRepository, dv validator.IDogValidator) IDogUsecase {
	return &dogUsecase{dr, dv}
}

func (du *dogUsecase) GetAllUsersDogs() ([]model.DogResponse, error) {
	dogs := []model.Dog{}
	if err := du.dr.GetAllUsersDogs(&dogs); err != nil {
		return nil, err
	}
	resDogs := []model.DogResponse{}
	for _, v := range dogs {
		t := model.DogResponse{
			ID:        v.ID,
			Img:       v.Img,
			Caption:   v.Caption,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			UserId:    v.UserId,
		}
		resDogs = append(resDogs, t)
	}
	return resDogs, nil
}

func (du *dogUsecase) GetAllDogs(userId uint) ([]model.DogResponse, error) {
	dogs := []model.Dog{}
	if err := du.dr.GetAllDogs(&dogs, userId); err != nil {
		return nil, err
	}
	resDogs := []model.DogResponse{}
	for _, v := range dogs {
		t := model.DogResponse{
			ID:        v.ID,
			Img:       v.Img,
			Caption:   v.Caption,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			UserId:    userId,
		}
		resDogs = append(resDogs, t)
	}
	return resDogs, nil
}

func (du *dogUsecase) GetDogById(dogId uint) (model.DogResponse, error) {
	dog := model.Dog{}
	if err := du.dr.GetDogById(&dog, dogId); err != nil {
		return model.DogResponse{}, err
	}
	resDog := model.DogResponse{
		ID:        dog.ID,
		Img:       dog.Img,
		Caption:   dog.Caption,
		CreatedAt: dog.CreatedAt,
		UpdatedAt: dog.UpdatedAt,
		UserId:    dog.UserId,
	}
	return resDog, nil
}

func (du *dogUsecase) CreateDog(dog model.Dog) (model.DogResponse, error) {
	//バリデーションを追加
	if err := du.dv.DogValidate(dog); err != nil {
		return model.DogResponse{}, err
	}

	if err := du.dr.CreateDog(&dog); err != nil {
		return model.DogResponse{}, err
	}
	resDog := model.DogResponse{
		ID:        dog.ID,
		Img:       dog.Img,
		CreatedAt: dog.CreatedAt,
		UpdatedAt: dog.UpdatedAt,
	}
	return resDog, nil
}

func (du *dogUsecase) UpdateDog(dog model.Dog, userId uint, dogId uint) (model.DogResponse, error) {
	//バリデーションを追加
	if err := du.dv.DogValidate(dog); err != nil {
		return model.DogResponse{}, err
	}

	if err := du.dr.UpdateDog(&dog, userId, dogId); err != nil {
		return model.DogResponse{}, err
	}
	resDog := model.DogResponse{
		ID:        dog.ID,
		Img:       dog.Img,
		CreatedAt: dog.CreatedAt,
		UpdatedAt: dog.UpdatedAt,
	}
	return resDog, nil
}

func (tu *dogUsecase) DeleteDog(userId uint, dogId uint) error {
	if err := tu.dr.DeleteDog(userId, dogId); err != nil {
		return err
	}
	return nil
}
