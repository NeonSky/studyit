package main

import (
	"github.com/NeonSky/studyit/backend/login"
	"github.com/NeonSky/studyit/backend/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	users.SetupRoutes(router)
	login.SetupRoutes(router)
	router.Run(":9000")
}
