package service

import (
	"demo1_gogin_api/db"
	"demo1_gogin_api/errno"
	"demo1_gogin_api/log"
	. "demo1_gogin_api/models"
	"demo1_gogin_api/redis"
	"demo1_gogin_api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
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
		log.UserLog.Error("Query from db error", err)
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
// @Success 200 {object} UserInfo "{"code":0,"message":"OK","data":{"xx":"xx"}}"
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
		log.UserLog.Error("Query from userinfo with Username failed", err)
		utils.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found")), nil)
		return
	}

	if err := db.SqlDB.Where("user_name = ?", userreq.Username).Where("user_passwd = ?", userreq.Password).First(&userinfo).Error; err != nil {
		utils.SendResponse(c, errno.New(errno.ErrPasswdIncorrect, fmt.Errorf("Password incorrect")), nil)
		return
	}

	token, err := utils.GenShortId()
	if err != nil {
		log.UserLog.Error("Generate uuid faied!", err)
		utils.SendResponse(c, errno.InternalServerError, nil)
	}
	log.UserLog.Info("Login with user[%s] success, token is [%s]", userreq.Username, token)

	if err := redis.Set(token, userinfo, time.Minute*time.Duration(viper.GetInt("token_expired"))); err != nil {
		log.UserLog.Error("Save token to redis failed!", err)
		utils.SendResponse(c, errno.InternalServerError, nil)
	}
	log.UserLog.Info("Save userinfo to redis success.")
	c.Header("X-Auth-Token", token)
	utils.SendResponse(c, nil, userinfo)
}
