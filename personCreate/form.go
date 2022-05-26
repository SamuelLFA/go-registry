package personcreate

import (
	personshare "github.com/SamuelLFA/goregister/personShare"
	"github.com/SamuelLFA/goregister/share/hash"
)

type PersonCreateForm struct {
	Name     string `form:"name" binding:"required,min=1,max=255"`
	Email    string `form:"email" binding:"required,min=1,max=255,unique,email"`
	Password string `form:"password" binding:"required,min=1,max=255"`
}

func toModel(form PersonCreateForm) personshare.PersonModel {
	pass, _ := hash.HashPassword(form.Password)
	return personshare.PersonModel{
		Name:     form.Name,
		Email:    form.Email,
		Password: pass,
	}
}
