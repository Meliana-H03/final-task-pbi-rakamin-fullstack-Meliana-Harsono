package main

import (
	"github.com/Meliana03/go-personalizeapi/controllers/photoscontroller"
	"github.com/Meliana03/go-personalizeapi/controllers/userscontroller"
	"github.com/Meliana03/go-personalizeapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/users/register", userscontroller.RegisterUser)
	r.GET("/users/login", userscontroller.LoginUser)
	r.GET("/users/validate", userscontroller.Validate)

	authorized := r.Group("/")
	authorized.Use(models.RequireAuth)
	{
		// Photos
		authorized.PUT("/users", userscontroller.UpdateUser)
		authorized.DELETE("/users", userscontroller.DeleteUser)

		authorized.POST("/photos", photoscontroller.CreatePhoto)
		authorized.GET("/photos", photoscontroller.GetPhotos)
		authorized.PUT("/photos/:photoId", photoscontroller.UpdatePhoto)
		authorized.DELETE("/photos/:photoId", photoscontroller.DeletePhoto)
	}

	r.Run()
}
