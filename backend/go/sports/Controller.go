package sports

import (
	"github.com/gin-gonic/gin"
)

func GetAllSports(c *gin.Context) {
	GetAllSportsDB(c)
}

func GetOneSport(c *gin.Context) {
	var slug string

	slug = c.Query("slug")

	GetOneSportDB(slug, c)
}

func CreateSport(c *gin.Context) {
	var sport SportModel

	c.BindJSON(&sport)

	CreateSportDB(&sport, c)
}

func UpdateSport(c *gin.Context) {
	var slug string
	var name string
	var img string

	slug = c.Query("slug")
	name = c.Query("name")
	img = c.Query("img")

	values := []string {slug, name, img}
	
	UpdateSportDB(values, c)

}

func DeleteSport(c *gin.Context) {
	var slug string

	slug = c.Query("slug")

	DeleteSportDB(slug, c)

} 