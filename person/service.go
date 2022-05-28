package person

import (
	"gorm.io/gorm"
)

type IPersonService interface {
	Save(model *PersonModel)
	Read(model *PersonModel, id int64)
}

type PersonService struct {
	DB *gorm.DB
}

func (p PersonService) Save(model *PersonModel) {
	p.DB.Save(model)
}

func (p PersonService) Read(model *PersonModel, id int64) {
	p.DB.First(&model, id)
}
