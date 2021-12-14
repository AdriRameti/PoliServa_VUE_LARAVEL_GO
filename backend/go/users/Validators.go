package users

import (
	"poliserva/Common"
	"github.com/gin-gonic/gin"
)

type UserWithRoleModelValidator struct {
	UUID string `form:"uuid" json:"uuid" binding:"required"`
	userWithRoleModel UserModel `json:"-"`
}

func NewUserWithRoleValidator() UserWithRoleModelValidator {
	return UserWithRoleModelValidator{}
}

func (user *UserWithRoleModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, user)

	if err != nil {
		return err
	}

	user.userWithRoleModel.UUID = user.UUID;

	return nil
}