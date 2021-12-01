package reservations

import (
	// "log"
	"poliserva/Common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReservationModelValidator struct{
		Id          uint   `json:"id" binding:"required"`
		Id_user     uint   `json:"id_user" binding:"required"`
		Id_court    uint   `json:"id_court" binding:"required"` 
		Date        string `json:"date" binding:"max=10,min=10"`
		Hini        string `json:"hini" binding:"max=5,min=5"`
		Hfin        string `json:"hfin" binding:"max=5,min=5"`
		Total_price int    `json:"total_price" binding:"min=1"`
	reservationModel ReservationModel `json:"-"`
}

func NewReservationModelValidator() ReservationModelValidator {
	return ReservationModelValidator{}
}

func (R *ReservationModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, R)
	if err != nil {
		return err
	}
	hini:= common.ValidateHour(R.Hini)
	hfin:= common.ValidateHour(R.Hfin)
	date:= common.ValidateDate(R.Date)
	if hini == false || hfin == false || date == false{
		c.AbortWithStatus(http.StatusNotFound);
	}else{
		R.reservationModel.Id = R.Id
		R.reservationModel.Id_user = R.Id_user
		R.reservationModel.Id_court = R.Id_court
		R.reservationModel.Date = R.Date
		R.reservationModel.Hini = R.Hini
		R.reservationModel.Hfin = R.Hfin
		R.reservationModel.Total_price = R.Total_price
	}

	return nil
}