package auth

import (
	"github.com/SamuelLFA/goregister/config/database"
	personshare "github.com/SamuelLFA/goregister/personShare"
	"github.com/SamuelLFA/goregister/share/hash"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
}

func StaticLoginService() LoginService {
	return &loginInformation{}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	var personData personshare.PersonModel
	database.DB.First(&personData, personshare.PersonModel{Email: email})

	if personData.Email == "" {
		return false
	}

	return personData.Email == email && hash.CheckPasswordHash(password, personData.Password)
}
