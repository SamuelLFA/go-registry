package validations

import (
	"github.com/SamuelLFA/goregister/config/database"
	p "github.com/SamuelLFA/goregister/person"
	"github.com/go-playground/validator/v10"
)

func Unique(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	var person p.PersonModel
	database.DB.First(&person, p.PersonModel{Email: email})

	return person.ID == 0
}
