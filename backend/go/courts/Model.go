package courts

import (
	"poliserva/Config"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CourtModel struct {
	Id	uint `json:"id"`
	Id_sport	uint	`gorm:"ForeignKey:ID" json:"id_sport"`
	Sector	int	`json:"sector"`
	Price_h	int	`json:"price_h"`
	Img string `json:"img"`
	Sports SportModel `gorm:"ForeignKey:Id_sport" json:"sports"`
}
type SportModel struct {
	ID	uint `json:"id"`
	Slug	string `json:"slug"`
	Name	string	`json:"name"`
	Img	string	`json:"img"`
}
func GetAllCourtsDB(c *gin.Context) []CourtModel {
	var courts []CourtModel
	err := Config.DB.Preload("Sports").Find(&courts).Error;
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}
	fmt.Println(courts)
	return courts
}
func GetDateOneReservation(MyArray []string, c *gin.Context) ([]CourtModel){
	var courts []CourtModel
	var err error
	Date:= MyArray[0]
	Hini:= MyArray[1]
	Hfin := MyArray[2]
	Slug := MyArray[3]

	if Slug != "undefined" {
		err = Config.DB.Preload("Sports").Raw("SELECT * FROM courts c WHERE c.id_sport = (SELECT s.id FROM sports s WHERE s.slug = ?) AND c.id NOT IN (SELECT r.id_court from reservations r where r.hini = ? and r.hfin = ? and r.date = ?)",Slug,Hini,Hfin,Date).Find(&courts).Error
	} else {
		err = Config.DB.Preload("Sports").Raw("SELECT * FROM courts c WHERE c.id NOT IN (SELECT r.id_court from reservations r  where r.hini = ? and r.hfin = ? and r.date = ?)",Hini,Hfin,Date).Find(&courts).Error
	}
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}
	
	return courts
}
func GetOneCourtDB(id string, c *gin.Context) CourtModel {
	var court CourtModel

	if err := Config.DB.Where("id = ?", id).Find(&court).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	return court
}

func GetCourtsBySportDB(slug string, c *gin.Context) []CourtModel {
	var courts []CourtModel

	if err := Config.DB.Preload("Sports").Raw("SELECT * FROM courts c WHERE c.id_sport = (SELECT id FROM sports s WHERE s.slug = ?)", slug).Find(&courts).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	return courts
}

func GetCourtsCarouselDB(c *gin.Context) []CourtModel {
	var courts []CourtModel

	if err := Config.DB.Limit(3).Find(&courts).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound);
		fmt.Println("Status:", err);
	}

	return courts
}

func CreateCourtDB(court *CourtModel, c *gin.Context) error {

	if err := Config.DB.Create(court).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return err
	} else {
		return nil
	}

}

func UpdateCourtDB(id string, court *CourtModel, c *gin.Context) (courtR CourtModel, errR error) {

	var courtF CourtModel

	if err := Config.DB.Where("id = ?", id).Find(&courtF).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
		return courtF, err
	} else {
		courtF.Id_sport = court.Id_sport
		courtF.Sector = court.Sector
		courtF.Price_h = court.Price_h
		Config.DB.Save(&courtF)
		return courtF, nil
	}

}

func DeleteCourtDB(id string, c *gin.Context) {

	var court CourtModel

	if errF := Config.DB.Where("id = ?", id).Delete(&court).Error; errF == nil {
		c.JSON(http.StatusOK, "Se ha eliminado con éxito")
	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (b *CourtModel) TableName() string {
	return "courts"
}
func (b *SportModel) TableName() string {
	return "sports"
}
