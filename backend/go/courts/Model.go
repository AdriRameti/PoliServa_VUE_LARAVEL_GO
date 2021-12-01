package courts

import (
	"poliserva/Config"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CourtModel struct {
	Id	uint `json:"id"`
	Id_sport	uint	`json:"id_sport"`
	Sector	int	`json:"sector"`
	Price_h	int	`json:"price_h"`
}

func GetAllCourtsDB(c *gin.Context) []CourtModel {
	var courts []CourtModel

	if err := Config.DB.Find(&courts).Error; err != nil {
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

func (b *CourtModel) TableName() string {
	return "courts"
}