package personread

import (
	personshare "github.com/SamuelLFA/goregister/personShare"
)

type PersonReadDTO struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func toDTO(model personshare.PersonModel) PersonReadDTO {
	return PersonReadDTO{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}
}
