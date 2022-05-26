package routes

import (
	"net/http"

	"github.com/SamuelLFA/goregister/auth"
	personcreate "github.com/SamuelLFA/goregister/personCreate"
	personread "github.com/SamuelLFA/goregister/personRead"
	"github.com/SamuelLFA/goregister/share/validations"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleRequest() {
	var loginService auth.LoginService = auth.StaticLoginService()
	var jwtService auth.JWTService = auth.JWTAuthService()
	var loginController auth.LoginController = auth.LoginHandler(loginService, jwtService)

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", validations.Unique)
	}

	r.POST("/people", personcreate.HandleRequest)
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
	r.GET("/people/:id", personread.HandleRequest)

	r.Run()
}
