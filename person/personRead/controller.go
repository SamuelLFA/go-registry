package personread

import (
	"net/http"
	"strconv"

	person "github.com/SamuelLFA/goregister/person"
	"github.com/gin-gonic/gin"
)

func ReadPersonHandle(service person.PersonService) PersonReadController {
	return PersonReadController{
		service: service,
	}
}

type PersonReadController struct {
	service person.IPersonService
}

func (p *PersonReadController) Read(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be an integer"})
		return
	}
	var model person.PersonModel
	p.service.Read(&model, id)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, toDTO(model))
}
