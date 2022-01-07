package reservations

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"poliserva/Config"
	courtsP "poliserva/courts"
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

func CreateNewReservation(reservations []string, c *gin.Context) bool {
	var reser []ReservationModel
	id_user:= reservations[0]
	id_court:= reservations[1]
	date:= reservations[2]
	hini:= reservations[3]
	hfin:= reservations[4]
	total_price:= reservations[5]
	err:= Config.DB.Raw("INSERT INTO reservations (id_user,id_court,date,hini,hfin,total_price) VALUES (?,?,?,?,?,?)",id_user,id_court,date,hini,hfin,total_price).Find(&reser).Error
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}
	return true
}
func GetDateOneReservation(MyArray []string, c *gin.Context) ([]courtsP.CourtModel){
	var courts []courtsP.CourtModel
	var err error
	Date:= MyArray[0]
	Hini:= MyArray[1]
	Hfin := MyArray[2]
	Slug := MyArray[3]

	if Slug != "undefined" {
		err = Config.DB.Model(&courts).Preload("Sports").Raw("SELECT * FROM courts c WHERE c.id_sport = (SELECT s.id FROM sports s WHERE s.slug = ?) AND c.id NOT IN (SELECT r.id_court from reservations r where r.hini = ? and r.hfin = ? and r.date = ?)",Slug,Hini,Hfin,Date).Find(&courts).Error
	} else {
		err = Config.DB.Raw("SELECT * FROM courts c WHERE c.id NOT IN (SELECT r.id_court from reservations r  where r.hini = ? and r.hfin = ? and r.date = ?)",Hini,Hfin,Date).Preload("Sports").Find(&courts).Error
	}

	if err != nil{
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	return courts
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
