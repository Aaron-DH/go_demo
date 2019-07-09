package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	db "demo1_gogin_api/databases"
	."demo1_gogin_api/common"
	."demo1_gogin_api/models"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsers(c *gin.Context) {
	var users []UserInfo
	if err := db.SqlDB.Find(&users).Error; err != nil {
		log.Error("Query from db error", err)
		SendResponse(c, DBError, nil)
	}

	SendResponse(c, nil, users)
}

func Login(c *gin.Context) {
	var userreq CreateRequest
	var userinfo UserInfo
	if err := c.BindJSON(&userreq); err != nil {
		SendResponse(c, ErrBind, nil)
		return
	}

	if userreq.Username == "" || userreq.Password == "" {
		SendResponse(c, ParamsError, nil)
		return
	}

	if err := db.SqlDB.Where("user_name = ?", userreq.Username).First(&userinfo).Error; err != nil {
		log.Error("Query from userinfo with Username failed", err)
		SendResponse(c, New(ErrUserNotFound, fmt.Errorf("username can not found")), nil)
		return
	}

	if err := db.SqlDB.Where("user_name = ?", userreq.Username).Where("user_passwd = ?", userreq.Password).First(&userinfo).Error; err != nil {
		SendResponse(c, New(ErrPasswdIncorrect, fmt.Errorf("Password incorrect")), nil)
		return
	}

	SendResponse(c, nil, userinfo)
}
