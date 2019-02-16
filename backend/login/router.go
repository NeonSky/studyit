package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/NeonSky/studyit/backend/users"
)

/* Public */
func SetupRoutes(router *gin.Engine) {
	router.GET("/login/:mail", loginUser)
}

/* GET */
func loginUser(c *gin.Context) {
	mail := c.Param("mail")
	fmt.Println("User with mail is trying to login: " + mail)
	user := users.RequestUser(mail)
	if user.Email != "" {
		c.JSON(http.StatusOK, user)
	} else {
		c.Status(404)
	}
}
