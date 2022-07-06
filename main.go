package main

import (
	"sesi12-final/controllers"
	"sesi12-final/db"
	"sesi12-final/middlewares"

	"github.com/gin-gonic/gin"
)

var PORT = ":1334"

func init() {
	db.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "API Running properly.")
	})
	r.POST("users/register", controllers.Register)
	r.POST("users/login", controllers.Login)

	usersRouter := r.Group("/users")
	{
		usersRouter.Use(middlewares.Validate())
		usersRouter.PUT("/:id", controllers.Update)
		usersRouter.DELETE("/", controllers.Delete)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Validate())
		photoRouter.POST("/", controllers.AddPost)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:id", controllers.UpdatePhoto)
		photoRouter.DELETE("/:id", controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Validate())
		commentRouter.POST("/", controllers.AddComment)
		commentRouter.PUT("/:id", controllers.UpdateComment)
		commentRouter.DELETE("/:id", controllers.DeleteComment)
	}

	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Validate())
		sosmedRouter.GET("/", controllers.GetSocialMedia)
		sosmedRouter.POST("/", controllers.AddSocialMedia)
		sosmedRouter.PUT("/:id", controllers.UpdateSocialMedia)
		sosmedRouter.DELETE("/:id", controllers.DeleteSocial)
	}

	r.Run(PORT)
}
