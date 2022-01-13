package users

import (
	"poliserva/Common"
	"github.com/gin-gonic/gin"
	// "net/http"
	// "fmt"
)

type UserModelValidator struct {
	// ID uint `form:"id" json:"id" binding:"required"`
	UUID string `form:"uuid" json:"uuid" binding:"omitempty"`
	Name string `form:"name" json:"name" binding:"required"`
	Surnames string `form:"surnames" json:"surnames" binding:"required"`
	Mail string `form:"mail" json:"mail" binding:"required"`
	Pass string `form:"pass" json:"pass" binding:"required"`
	Img string `form:"img" json:"img" binding:"omitempty"`
	Role string `form:"role" json:"role" binding:"required"`
	// Google2fa_secret string `form:"" json:"" binding:"omitempty"`
	Is_blocked bool `form:"is_blocked" json:"is_blocked" binding:"omitempty"`
	userModel UserModel `json:"-"`
}

func NewUserValidator() UserModelValidator {
	return UserModelValidator{}
}

func (user *UserModelValidator) Bind(c *gin.Context) (error, string) {
	err := common.Bind(c, user)

	if (c.Query("uuid") != "") {
		user.UUID = c.Query("uuid")
	}

	if err != nil {
		return err, "can't bind"
	}

	if errFM := CheckMail(user, c); errFM == true {
		return nil, "already used"
	} else {

		if (user.UUID != "") {
			user.userModel.UUID = user.UUID
		} else {
			user.userModel.UUID = common.GenerateUUID()
		}

		if (user.Pass != "") {
			user.userModel.Pass, _ = common.HashPass(user.Pass)
		}

		if (user.Img != "") {
			user.userModel.Img = user.Img
		} else {
			user.userModel.Img = "https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/640px-User_icon_2.svg.png"
		}

		user.userModel.Name = user.Name
		user.userModel.Surnames = user.Surnames
		user.userModel.Mail = user.Mail
		user.userModel.Role = user.Role
		user.userModel.Is_blocked = user.Is_blocked
	}
	
	return nil, "not used"
}
