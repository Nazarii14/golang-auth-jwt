package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	var err error
	userType := c.GetString("user_type")

	if userType != role {
		return errors.New("Unauthorized")
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	var err error
	if userType == "USER" && userId != uid {
		err = errors.New("Unauthorized")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
