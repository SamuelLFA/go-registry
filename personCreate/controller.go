package personcreate

import (
	"net/http"

	database "github.com/SamuelLFA/goregister/config/database"
	personshare "github.com/SamuelLFA/goregister/personShare"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleRequest(c *gin.Context) {
	var form PersonCreateForm
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var model personshare.PersonModel = toModel(form)
	database.DB.Save(&model)
	c.JSON(201, model)
}
