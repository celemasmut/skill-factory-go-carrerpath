package app

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/controllers"
	"github.com/avalith-net/skill-factory-go-carrerpath/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Career Path API
// @version 1.0
// @description This is an application that allows you to manage your career path
// @termsOfService http://swagger.io/terms/

// @contact.name Career Path API Support
// @contact.url http://www.swagger.io/support
// @contact.email careerpath.avalith@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func mapUrls() {
	router.Use(middlewares.CheckDB())
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/userpath", middlewares.ValidateJWT(), controllers.ShowUserPath)
	router.POST("/passwordRecovery", controllers.PasswordRecovery)

	//use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}