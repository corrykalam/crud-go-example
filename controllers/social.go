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
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

func AddSocialMedia(c *gin.Context) {
	req := models.Social{}
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
		"id":               req.ID,
		"name":             req.Name,
		"social_media_url": req.SocialMediaUrl,
		"user_id":          req.UserId,
		"created_at":       req.CreatedAt,
	})
}

func GetSocialMedia(c *gin.Context) {
	var result []map[string]interface{}
	// err = db.GetDB().Debug().Create(&req).Error

	err := db.GetDB().Debug().Raw("SELECT *  FROM socials INNER JOIN users ON socials.user_id = users.id;").Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i < len(result); i++ {
		delete(result[i], "password")
	}
	c.JSON(201, gin.H{
		"socialmedias": result,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	req := models.Social{}
	dbRes := models.Social{}
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
	err = db.GetDB().Exec("UPDATE socials SET name = ?, social_media_url = ? WHERE id = ?", req.Name, req.SocialMediaUrl, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               dbRes.ID,
		"name":             req.Name,
		"social_media_url": req.SocialMediaUrl,
		"user_id":          dbRes.UserId,
		"updated_at":       dbRes.UpdatedAt,
	})
}

func DeleteSocial(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = db.GetDB().Exec("DELETE FROM socials where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
