package common
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"regexp"
	"log"
)
func Bind(c *gin.Context, obj interface{}) error{
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

func ValidateHour(hour string) bool{
	match, _ := regexp.MatchString("[0-9][0-9]:[0-9][0-9]", hour)
	return match
}
func ValidateDate(date string) bool{
	match,_:= regexp.MatchString("[0-9]{2}/[0-9]{2}/[0-9]{4}", date)
	log.Println(match)
	return match
}