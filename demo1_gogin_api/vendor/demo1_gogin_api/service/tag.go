package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	."demo1_gogin_api/models"
	db "demo1_gogin_api/databases"
	"time"
)

func GetTags(c *gin.Context) {
	var tags []TbTag
	if err := db.SqlDB.Find(&tags).Error; err != nil {
		c.AbortWithStatus(404)
		log.Error("Query from db error", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : tags,
	})
}

func UpdateTag(c *gin.Context) {
	var tag TbTag
	id := c.Params.ByName("id")

	log.Infof("Begin update tag: %s", id)
	if err := db.SqlDB.Where("tag_id = ?", id).First(&tag).Error; err != nil {
		log.Error("Query from tag with id failed", err)
		c.JSON(404, "Update failed, Tagid:" + id + " Not Found")
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
		log.Error("Insert into db failed", err)
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, tag)
}
