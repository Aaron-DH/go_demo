package service

import (
	"demo1_gogin_api/db"
	"demo1_gogin_api/log"
	. "demo1_gogin_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetTags(c *gin.Context) {
	var tags []TbTag
	if err := db.SqlDB.Find(&tags).Error; err != nil {
		c.AbortWithStatus(404)
		log.UserLog.Error("Query tags from db error.", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tags,
	})
}

func UpdateTag(c *gin.Context) {
	var tag TbTag
	id := c.Params.ByName("id")

	log.UserLog.Info("Begin update tag: %s", id)
	if err := db.SqlDB.Where("tag_id = ?", id).First(&tag).Error; err != nil {
		log.UserLog.Error("Query tag from db with id failed.", err)
		c.JSON(404, "Update failed, Tagid:"+id+" Not Found")
		return
	}

	c.BindJSON(&tag)

	db.SqlDB.Save(&tag)

	c.JSON(http.StatusOK, tag)
}

func CreateTag(c *gin.Context) {
	var tag TbTag
	c.BindJSON(&tag)
	tag.CreateTime = time.Now()

	if err := db.SqlDB.Create(&tag).Error; err != nil {
		log.UserLog.Error("Insert taginfo to db failed.", err)
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, tag)
}
