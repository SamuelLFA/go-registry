package personcreate

import (
	person "github.com/SamuelLFA/goregister/person"
	"github.com/SamuelLFA/goregister/share/hash"
)

type PersonCreateForm struct {
	Name     string `form:"name" binding:"required,min=1,max=255"`
	Email    string `form:"email" binding:"required,min=1,max=255,unique,email"`
	Password string `form:"password" binding:"required,min=1,max=255"`
}

func (p *PersonCreateForm) ToModel() person.PersonModel {
	pass, _ := hash.HashPassword(p.Password)
	return person.PersonModel{
		Name:     p.Name,
		Email:    p.Email,
		Password: pass,
	}
}
