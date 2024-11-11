package routers

import (
	"swagger/controllers"
	"swagger/middlewares"

	"github.com/gin-gonic/gin"

	_ "swagger/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

// @title Cars API
// @version 2.0
// @description This is an application for activity cars API
// @host localhost:8080
// @BasePath /

func StartRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	carsRouter := r.Group("/api")
	{
		carsRouter.Use(middlewares.Authentication())
		carsRouter.GET("/cars", controllers.GetAllCars)
		carsRouter.POST("/cars", controllers.CreateCars)
	}

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	return r
}
