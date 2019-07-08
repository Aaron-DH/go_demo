package tag

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	."goapi/models"
)

func GetTags(c *gin.Context) {
	var t Tag
	tags, err := t.GetTags()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : tags,
	})
}