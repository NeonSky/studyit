package main

import (
	"github.com/gin-gonic/gin"
	"github.com/studyit/backend/login"
	"github.com/studyit/backend/users"
)

func main() {
	router := gin.Default()
	users.SetupRoutes(router)
	login.SetupRoutes(router)
	router.Run(":9000")
}
