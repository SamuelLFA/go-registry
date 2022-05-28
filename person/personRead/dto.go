package personread

import (
	person "github.com/SamuelLFA/goregister/person"
)

type PersonReadDTO struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func toDTO(model person.PersonModel) PersonReadDTO {
	return PersonReadDTO{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}
}
