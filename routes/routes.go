package routes

import (
	"net/http"

	"github.com/SamuelLFA/goregister/auth"
	"github.com/SamuelLFA/goregister/config/database"
	person "github.com/SamuelLFA/goregister/person"
	personcreate "github.com/SamuelLFA/goregister/person/personCreate"
	personread "github.com/SamuelLFA/goregister/person/personRead"
	"github.com/SamuelLFA/goregister/share/validations"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleRequest() {
	var loginService auth.LoginService = auth.StaticLoginService()
	var personService person.PersonService = person.PersonService{DB: database.DB}
	var jwtService auth.JWTService = auth.JWTAuthService()
	var loginController auth.LoginController = auth.LoginHandler(loginService, jwtService)
	var createPersonController personcreate.PersonCreateController = personcreate.CreatePersonHandle(personService)
	var personReadController personread.PersonReadController = personread.ReadPersonHandle(personService)

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", validations.Unique)
	}

	r.POST("/people", createPersonController.Create)
	r.POST("/auth", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	r.Use(auth.AuthorizeJWT())
	r.GET("/people/:id", personReadController.Read)

	r.Run()
}
