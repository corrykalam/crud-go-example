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

func AddPost(c *gin.Context) {
	req := models.Photo{}
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
		"title":      req.Title,
		"caption":    req.Caption,
		"photo_url":  req.PhotoUrl,
		"user_id":    req.UserId,
		"created_at": req.CreatedAt,
	})
}

func GetPhoto(c *gin.Context) {
	var result []map[string]interface{}

	err := db.GetDB().Debug().Raw("SELECT photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email FROM photos INNER JOIN users ON photos.user_id = users.id;").Scan(&result).Error
	fmt.Println(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdatePhoto(c *gin.Context) {
	req := models.Photo{}
	dbRes := models.Photo{}
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
	err = db.GetDB().Debug().Where("id = ?", id).Last(&dbRes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.GetDB().Exec("UPDATE photos SET title = ?, caption = ?, photo_url = ? WHERE id = ?", req.Title, req.Caption, req.PhotoUrl, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         dbRes.ID,
		"title":      req.Title,
		"caption":    req.Caption,
		"photo_url":  req.PhotoUrl,
		"user_id":    dbRes.UserId,
		"updated_at": dbRes.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = db.GetDB().Exec("DELETE FROM photos where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
