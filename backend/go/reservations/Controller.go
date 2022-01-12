package reservations

import (
	"log"
	"net/http"
	"poliserva/Common"
	courtsP "poliserva/courts"

	"github.com/gin-gonic/gin"
)

func GetAllReservations(c *gin.Context) {
	var reservations []ReservationModel
	consult:=GetAllReservation(&reservations,c)
	serializer:= ReservationsSerializer{c, *consult}
	c.JSON(http.StatusCreated,serializer.Response())
}
func GetUserReservation(c *gin.Context){
	id:= c.Query("id")
	fIni:= c.Query("fIni")
	fFin:= c.Query("fFin")
	arr:= []string {id,fIni,fFin}
	consult:= GetUserReservations(arr,c)
	serializer:= ReservationsSerializer{c, consult}
	c.JSON(http.StatusCreated,serializer.Response())
}
func GetUserSportReservation(c *gin.Context){
	id:= c.Query("id")
	dashboard:= c.Query("type")
	consult:=GetUserSportsReservation(id,dashboard,c)
	c.JSON(http.StatusCreated,consult)
}
func GetUserCourtReservation(c *gin.Context){
	consult:=GetCourtsReservation(c)
	c.JSON(http.StatusCreated,consult)
}
func GetDateReservations(c *gin.Context){
	// var reservations []courtsP.CourtModel
	Date:= c.Query("date")
	hini:= c.Query("hini")
	hfin:= c.Query("hfin")
	slug:= c.Query("slug")
	if len(Date)<1||  len(hini)<1 || len(hfin)<1 || len(slug)<1 {
		return
	}
	validDate:= common.ValidateDate(Date)
	validHini:= common.ValidateHour(hini)
	validHfin:= common.ValidateHour(hfin)
	if validDate == false || validHini == false || validHfin == false{
		c.AbortWithStatus(http.StatusNotFound);
	}
	MyArray:= []string {Date,hini,hfin,slug}
	courts:=GetDateOneReservation(MyArray,c)
	serializer:= courtsP.CourtsSerializer{c, courts}
	c.JSON(http.StatusCreated, serializer.Response())
}
func CreateReservations(c *gin.Context){
	id_user:= c.GetHeader("id_user")
	id_court:= c.GetHeader("id_court")
	date:= c.GetHeader("date")
	hini:= c.GetHeader("hini")
	hfin:= c.GetHeader("hfin")
	total_price:= c.GetHeader("total_price")
	arr:= []string{id_user,id_court,date,hini,hfin,total_price}
	request:=CreateNewReservation(arr,c)
	c.JSON(http.StatusCreated,request)

}

func UpdateReservations(c *gin.Context){
	var reservations ReservationModel
	id := c.Params.ByName("id")
	if len(id)<1{
		log.Println("Id is missing")
		return
	}
	UpdateReservation(&reservations,id,c)
}

func DeleteReservations(c *gin.Context){
	var reservations ReservationModel
	id := c.Params.ByName("id")
	if len(id)<1{
		log.Println("Id is missing")
		return
	}
	DeleteReservation(&reservations,id,c)
}
