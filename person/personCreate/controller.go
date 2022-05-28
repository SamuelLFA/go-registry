package personcreate

import (
	"net/http"

	"github.com/SamuelLFA/goregister/person"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreatePersonHandle(service person.PersonService) PersonCreateController {
	return PersonCreateController{
		service: service,
	}
}

type PersonCreateController struct {
	service person.IPersonService
}

func (p *PersonCreateController) Create(c *gin.Context) {
	var form PersonCreateForm
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var model person.PersonModel = form.ToModel()
	p.service.Save(&model)
	c.JSON(201, model)
}
