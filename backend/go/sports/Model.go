package sports

import (
	"poliserva/Config"
	"fmt"
	"net/http"
	"math/rand"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

type SportModel struct {
	Id	uint `json:"id"`
	Slug	string `json:"slug"`
	Name	string	`json:"name"`
	Img	string	`json:"img"`
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

func GetAllSportsDB(c *gin.Context) {

	var sports []SportModel

	if err := Config.DB.Find(&sports).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	c.JSON(http.StatusOK, sports)

}

func GetOneSportDB(slug string,c *gin.Context) {

	var sport SportModel

	if err := Config.DB.Where("slug = ?", slug).Find(&sport).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, sport)
	}

}

func CreateSportDB(sport *SportModel, c *gin.Context) {

	if err := Config.DB.Create(sport).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, sport)
	}

}

func UpdateSportDB(values []string, c *gin.Context) {

	var sport SportModel
	var slug string
	var name string
	var img string

	slug = values[0]
	name = values[1]
	img = values[2]

	if errF := Config.DB.Where("slug = ?", slug).Find(&sport).Error; errF == nil {

		if name != "" {
			sport.Name = name
		}

		if img != "" {
			sport.Img = img
		}

		Config.DB.Save(&sport)

		c.JSON(http.StatusOK, sport)
	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

}

func DeleteSportDB(slug string, c *gin.Context) {

	var sportF SportModel

	if errF := Config.DB.Where("slug = ?", slug).Delete(&sportF).Error; errF == nil {
		c.JSON(http.StatusOK, "Se ha eliminado con éxito")
	} else {
		fmt.Println(errF.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

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