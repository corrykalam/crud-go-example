package controllers

import (
	"fmt"
	"net/http"
	"sesi12-final/db"
	"sesi12-final/helpers"
	"sesi12-final/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Register(c *gin.Context) {
	req := models.User{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := validate.Struct(req)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.GetDB().Debug().Create(&req).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":       req.ID,
		"username": req.Username,
		"email":    req.Email,
		"age":      req.Age,
	})
}

func Login(c *gin.Context) {
	req := models.User{}
	dbRes := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.GetDB().Debug().Where("username = ?", req.Username).Last(&dbRes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbRes.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid password"})
		return
	}
	token := helpers.GenerateToken(req.Username, dbRes.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Update(c *gin.Context) {
	req := models.User{}
	dbRes := models.User{}
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
	err = db.GetDB().Exec("UPDATE users SET email = ?, username = ? WHERE id = ?", req.Email, req.Username, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       dbRes.ID,
		"username": req.Username,
		"email":    req.Email,
		"age":      dbRes.Age,
	})
}

func Delete(c *gin.Context) {
	jwt := c.Request.Header.Get("Authorization")
	jwtParse, _ := helpers.VerifyToken(strings.Split(jwt, " ")[1])
	fmt.Println(jwtParse["id"])
	err := db.GetDB().Exec("DELETE FROM users where id = ?", jwtParse["id"]).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
