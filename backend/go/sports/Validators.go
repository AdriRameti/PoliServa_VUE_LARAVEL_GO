package sports

import (
	"poliserva/Common"
	"github.com/gin-gonic/gin"
)

type SportModelValidator struct {
	Name	string	`form:"name" json:"name" binding:"required"`
	Img	string	`form:"img" json:"img" binding:"required"`
	sportModel SportModel `json:"-"`
}

func NewModelValidator() SportModelValidator {
	return SportModelValidator{}
}

func (sport *SportModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, sport)

	if err != nil {
		return err
	}

	sport.sportModel.Name = sport.Name
	sport.sportModel.Img = sport.Img

	return nil
}