package courts

import (
	"poliserva/Common"
	"github.com/gin-gonic/gin"
)

type CourtModelValidator struct {
	Id_sport	uint	`form:"id_sport" json:"id_sport" binding:"required"`
	Sector	int	`form:"sector" json:"sector" binding:"required"`
	Price_h	int	`form:"price_h" json:"price_h" binding:"required"`
	Img string `form:"img" json:"img" binding:"required"`
	courtModel CourtModel `json:"-"`
}

func NewModelValidator() CourtModelValidator {
	return CourtModelValidator{}
}

func (court *CourtModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, court)

	if err != nil {
		return err
	}

	court.courtModel.Id_sport = court.Id_sport
	court.courtModel.Sector = court.Sector
	court.courtModel.Price_h = court.Price_h
	court.courtModel.Img = court.Img

	return nil
}












































// func ValidateCourt(court *CourtModel) bool {

	// var sport []sports.SportModel

	// fmt.Println(court.Id_sport)

	// if errF := Config.DB.Find(&sport).Error; errF != nil {
	// 	fmt.Println(sport),
	// 	return false
	// }

	// fmt.Println(sport)

	// return true
// }