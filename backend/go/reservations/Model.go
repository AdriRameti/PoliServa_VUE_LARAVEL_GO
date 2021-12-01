package reservations

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"poliserva/Config"
)

type ReservationModel struct {
	Id          uint   `json:"id"`
	Id_user     uint   `json:"id_user"`
	Id_court    uint   `json:"id_court"`
	Date        string `json:"date"`
	Hini        string `json:"hini"`
	Hfin        string `json:"hfin"`
	Total_price int    `json:"total_price"`
}

func GetAllReservation(reservations *[]ReservationModel, c *gin.Context) (*[]ReservationModel){
	if err := Config.DB.Find(&reservations).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	return reservations
}

func CreateNewReservation(reservations *ReservationModel, c *gin.Context) {
	if err:= Config.DB.Create(&reservations).Error; err!=nil{
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}
	// c.JSON(http.StatusAccepted,reservations)
}
func GetDateOneReservation(MyArray []string,reservations *[]ReservationModel, c *gin.Context) (*[]ReservationModel){
	Date:= MyArray[0]
	Hour:= MyArray[1]
	err:= Config.DB.Raw("Select * from reservations where not exists (select * from reservations r where ? between r.hini and r.hfin and r.date = ?)",Hour,Date).Find(&reservations).Error
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}
	return reservations
}

func UpdateReservation(reservations *ReservationModel,id string,c *gin.Context){

	if err:= Config.DB.Where("id = ?", id).Find(&reservations).Error;err != nil {
		c.JSON(http.StatusNotFound, reservations)
	}
	c.BindJSON(reservations)
	err := UpdateRes(reservations,id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reservations)
	}
}

func UpdateRes(reservations *ReservationModel, id string) (err error) {
	fmt.Println(reservations)
	Config.DB.Save(reservations)
	return nil
}

func DeleteReservation(reservations *ReservationModel,id string,c *gin.Context){
	if err:= Config.DB.Where("id = ?", id).Delete(reservations).Error; err!= nil{
		c.JSON(http.StatusNotFound, reservations)
	}
	c.JSON(http.StatusOK, reservations)
}

func (b *ReservationModel) TableName() string {
	return "reservations"
}
