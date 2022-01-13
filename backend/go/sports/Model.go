package sports

import (
	"poliserva/Config"
	"fmt"
	"net/http"
	"math/rand"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	courtsP "poliserva/courts"
)

type SportModel struct {
	ID	uint `json:"id"`
	Slug	string `json:"slug"`
	Name	string	`json:"name"`
	Img	string	`json:"img"`
	Courts []courtsP.CourtModel `gorm:"ForeignKey:Id_sport" json:"courts"`
}

func (sport *SportModel) BeforeCreate() {
	
	var sportF SportModel

	for {
		slugify(sport)

		if errF := Config.DB.Where("slug = ?", sport.Slug).Find(&sportF).Error; errF != nil {
			break
		}
		
	}

}

func (sport *SportModel) BeforeSave() {
	
	var sportF SportModel

	for {
		slugify(sport)

		if errF := Config.DB.Where("slug = ?", sport.Slug).Find(&sportF).Error; errF != nil {
			break
		}
		
	}

}

func GetAllSportsDB(c *gin.Context) []SportModel {

	var sports []SportModel

	if err := Config.DB.Find(&sports).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return sports

} 
// func GetIdSportDB(id,c *gin.Context)SportModel{
// 	var sport SportModel
// 	if err := Config.DB.Where("id = ?", id).Find(&sport).Error; err != nil {
// 		fmt.Println(err.Error())
// 		c.AbortWithStatus(http.StatusNotFound)
// 		return sport
// 	} else {
// 		return sport
// 	}
// }
func GetOneSportDB(slug string,c *gin.Context) SportModel {

	var sport SportModel

	if err := Config.DB.Where("slug = ?", slug).Find(&sport).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return sport
	} else {
		return sport
	}

}

func CreateSportDB(sport *SportModel, c *gin.Context) (sportR SportModel, errR error) {
	
	var sportRet SportModel

	if err := Config.DB.Create(sport).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		
		return sportRet, err
	} else {
		Config.DB.Where("slug = ?", sport.Slug).Find(&sportRet)
		return sportRet, err
	}

}

func UpdateSportDB(slug string, sport *SportModel, c *gin.Context) (sportR SportModel, errR error) {

	var sportF SportModel

	if errF := Config.DB.Where("slug = ?", slug).Find(&sportF).Error; errF == nil {

		sportF.Name = sport.Name
		sportF.Img = sport.Img

		Config.DB.Save(&sportF)

		return sportF, errF
	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return sportF, errF
	}

}

func DeleteSportDB(slug string, c *gin.Context) []SportModel {

	var sportF SportModel
	var courts []courtsP.CourtModel
	var sports []SportModel

	if errF := Config.DB.Where("slug = ?", slug).Preload("Courts").Find(&sportF).Error; errF == nil {

		Config.DB.Model(&sportF).Association("Courts").Find(&courts)

		if len(sportF.Courts) != 0 {

			if Config.DB.Model(&sportF).Association("Courts").Delete(sportF.Courts).Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Unable to delete courts association",})
			} else {
				Config.DB.Where("slug = ?", slug).Delete(&sportF)
	
				for _, court := range courts {
	
					courtsP.DeleteCourtDB(strconv.FormatUint(uint64(court.Id), 10), c)
				}

				if err := Config.DB.Find(&sports).Error; err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					fmt.Println("Status:", err)
				}

				return sports
			}
		} else {
			if Config.DB.Where("slug = ?", slug).Delete(&sportF).Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Unable to delete sports",})
			}

			if err := Config.DB.Find(&sports).Error; err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				fmt.Println("Status:", err)
			}

			return sports
		}

	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}

	return sports

}

func slugify(sport *SportModel) *SportModel {
	var sb strings.Builder

	sb.WriteString(sport.Name)
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(rand.Intn(9999)))
	sport.Slug = sb.String()

	return sport
}


func (b *SportModel) TableName() string {
	return "sports"
}