package controller

import (
	"log"
	"os"

	"github.com/avalith-net/skill-factory-go-carrerpath/middlewares"
	"github.com/avalith-net/skill-factory-go-carrerpath/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//LaunchingServer seteo mi puerto, el handler y pongo a escuchar al servidor
func LaunchingServer() {
	router := gin.Default()
	router.POST("/register", middlewares.CheckDB(), routers.Register)
	router.POST("/login", middlewares.CheckDB(), routers.Login)
	router.GET("/userpath", middlewares.CheckDB(), middlewares.ValidateJWT(), routers.ShowUserPath)
	router.POST("/passwordRecovery", middlewares.CheckDB(), routers.PasswordRecovery) //no esta creada

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Use(cors.Default())
	log.Fatal(router.Run())
}
