package validations

import (
	"github.com/SamuelLFA/goregister/config/database"
	personshare "github.com/SamuelLFA/goregister/personShare"
	"github.com/go-playground/validator/v10"
)

func Unique(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	var person personshare.PersonModel
	database.DB.First(&person, personshare.PersonModel{Email: email})

	return person.ID == 0
}
