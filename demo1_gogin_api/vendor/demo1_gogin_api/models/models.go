package models

import (
	db "demo1_gogin_api/databases"
	"time"
)

type CommonField struct {
	CreateID int `json:"create_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateID int `json:"update_id"`
	UpdateTime time.Time `json:"update_time"`
	IsDel int `json:"is_del"`
}

type TbTag struct {
	TagID int `json:"tag_id"`
	TagName string `json:"tag_name"`
	GroupID int `json:"group_id"`
	CreateID int `json:"create_id"`
	CreateTime time.Time `json:"create_time"`
}

//type MenuResource struct {
//	RescID int `json:"resc_id"`
//	RescName string `json:"resc_name"`
//	RescUrl string `json:"resc_url"`
//	RescIcon string `json:"resc_icon"`
//	CommonField
//}

type UserInfo struct {
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	UserPasswd string `json:"user_passwd"`
	UserStatus int `json:"user_status"`
	UserMail string `json:"user_mail"`
	CommonField
}

func Init() {
	db.SqlDB.AutoMigrate(&TbTag{}, &UserInfo{})
}
