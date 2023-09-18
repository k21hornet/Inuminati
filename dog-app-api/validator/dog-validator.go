package validator

import (
	"dog-app-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IDogValidator interface {
	DogValidate(dog model.Dog) error
}

type dogValidator struct{}

func NewDogValidator() IDogValidator {
	return &dogValidator{}
}

func (dv *dogValidator) DogValidate(dog model.Dog) error {
	return validation.ValidateStruct(&dog,
		validation.Field(
			&dog.Img,
			validation.Required.Error("img is required"),
		),
	)
}
