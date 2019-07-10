package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	db "demo1_gogin_api/databases"
	"demo1_gogin_api/utils"
	"demo1_gogin_api/errno"
	."demo1_gogin_api/models"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary List all users from database
// @Description Get Users
// @Tags UserInfo
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserInfo "{"code":0,"message":"OK","data":{"user_name":"kong","xx":"xx"}}"
// @Router /v1/users [get]
func GetUsers(c *gin.Context) {
	var users []UserInfo
	if err := db.SqlDB.Find(&users).Error; err != nil {
		log.Error("Query from db error", err)
		utils.SendResponse(c, errno.DBError, nil)
	}

	utils.SendResponse(c, nil, users)
}

// @Summary Login with username and password
// @Description Login
// @Tags Login
// @Accept  json
// @Produce  json
// @Param username formData string true "用户名"
// @Param password formData string true "用户密码"
// @Success 200 {object} models.UserInfo "{"code":0,"message":"OK","data":{"user_name":"kong","xx":"xx"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	var userreq CreateRequest
	var userinfo UserInfo
	if err := c.BindJSON(&userreq); err != nil {
		utils.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if userreq.Username == "" || userreq.Password == "" {
		utils.SendResponse(c, errno.ParamsError, nil)
		return
	}

	if err := db.SqlDB.Where("user_name = ?", userreq.Username).First(&userinfo).Error; err != nil {
		log.Error("Query from userinfo with Username failed", err)
		utils.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found")), nil)
		return
	}

	if err := db.SqlDB.Where("user_name = ?", userreq.Username).Where("user_passwd = ?", userreq.Password).First(&userinfo).Error; err != nil {
		utils.SendResponse(c, errno.New(errno.ErrPasswdIncorrect, fmt.Errorf("Password incorrect")), nil)
		return
	}

	utils.SendResponse(c, nil, userinfo)
}
