package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* Public */
func SetupRoutes(router *gin.Engine) {
	router.GET("/users/:mail", getUserByMail)
	router.GET("/users/", getUsers)
	router.POST("/users/", createUser)
	router.PUT("/users/:mail", updateUserByMail)
	router.DELETE("/users/:mail", deleteUserByMail)
}

/* GET */
func getUserByMail(c *gin.Context) {
	mail := c.Param("mail")
	user := RequestUser(mail)
	c.JSON(http.StatusOK, user)
}

func getUsers(c *gin.Context) {
	users := requestUsers()
	c.JSON(http.StatusOK, users)
}

/* POST */
func createUser(c *gin.Context) {
	var newUser User
	json.NewDecoder(c.Request.Body).Decode(&newUser)
	requestCreateUser(newUser)
	c.Status(http.StatusOK)
}

/* PUT */
func updateUserByMail(c *gin.Context) {
	mail := c.Param("mail")
	prevUser := RequestUser(mail)

	var newUser User
	json.NewDecoder(c.Request.Body).Decode(&newUser)

	requestUpdateUser(prevUser, newUser)
	c.JSON(http.StatusOK, newUser)
}

/* DELETE */
func deleteUserByMail(c *gin.Context) {
	mail := c.Param("mail")
	user := RequestUser(mail)
	requestDeleteUser(user)
	c.JSON(http.StatusOK, user)
}
