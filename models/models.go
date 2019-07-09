package models

import (
	db "go_api_demo/databases"
	"time"
)

type TbTag struct {
	TagID int `json:"tag_id"`
	TagName string `json:"tag_name"`
	GroupID int `json:"group_id"`
	CreateID int `json:"create_id"`
	CreateTime time.Time `json:"create_time"`
}

func Init() {
	db.SqlDB.AutoMigrate(&TbTag{})
}
