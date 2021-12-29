package common
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"regexp"
	"log"
	"crypto/rand"
  	"encoding/hex"
	"strings"
	"golang.org/x/crypto/bcrypt"
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

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
	  return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateUUID() string {

	var sb strings.Builder

	val1, _ := randomHex(4)
	val2, _ := randomHex(2)
	val3, _ := randomHex(2)
	val4, _ := randomHex(2)
	val5, _ := randomHex(6)

	sb.WriteString(val1)
	sb.WriteString("-")
	sb.WriteString(val2)
	sb.WriteString("-")
	sb.WriteString(val3)
	sb.WriteString("-")
	sb.WriteString(val4)
	sb.WriteString("-")
	sb.WriteString(val5)

	return sb.String()
}

func HashPass(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}