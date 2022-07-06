package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"sesi12-final/db"
	"sesi12-final/helpers"
	"sesi12-final/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	req := models.Comment{}
	jwt := c.Request.Header.Get("Authorization")
	jwtParse, _ := helpers.VerifyToken(strings.Split(jwt, " ")[1])
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(reflect.TypeOf(jwtParse["id"]))
	req.UserId = int(jwtParse["id"].(float64))

	err := db.GetDB().Create(&req).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         req.ID,
		"message":    req.Message,
		"photo_id":   req.PhotoId,
		"user_id":    req.UserId,
		"created_at": req.CreatedAt,
	})
}

// func GetPhoto(c *gin.Context) {
// 	var result []map[string]interface{}

// 	err := db.GetDB().Debug().Raw("SELECT photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email FROM photos INNER JOIN users ON photos.user_id = users.id;").Scan(&result).Error
// 	fmt.Println(result)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}

// 	c.JSON(http.StatusOK, result)
// }

func UpdateComment(c *gin.Context) {
	req := models.Comment{}
	var dbRes map[string]interface{}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = db.GetDB().Debug().Raw("SELECT photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.updated_at FROM photos INNER JOIN comments ON photos.user_id = comments.user_id WHERE comments.id = ? LIMIT 1;", id).Last(&dbRes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.GetDB().Exec("UPDATE comments SET message = ? WHERE id = ?", req.Message, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         dbRes["id"],
		"title":      dbRes["title"],
		"caption":    dbRes["caption"],
		"photo_url":  dbRes["photo_url"],
		"user_id":    dbRes["user_id"],
		"updated_at": dbRes["updated_at"],
	})
}

func DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = db.GetDB().Exec("DELETE FROM comments where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
