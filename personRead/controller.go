package personread

import (
	"net/http"
	"strconv"

	database "github.com/SamuelLFA/goregister/config/database"
	personshare "github.com/SamuelLFA/goregister/personShare"
	"github.com/gin-gonic/gin"
)

func HandleRequest(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be an integer"})
		return
	}
	var model personshare.PersonModel
	database.DB.First(&model, id)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, toDTO(model))
}
