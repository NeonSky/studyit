package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var state string

/* Public */
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginHandler(c *gin.Context) {
	state = RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	c.Writer.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
}

func StudyHandler(c *gin.Context) {
	fmt.Println("Reached study handler")
	courseID := c.Param("id")
	fmt.Println(courseID)

	session := sessions.Default(c)
	userID := session.Get("user-id")
	fmt.Println("...and our user-id is", userID)

	c.HTML(http.StatusOK, "index.html", nil)
}

func UserHandler(c *gin.Context) {
	fmt.Println("Reached user handler")
	courseID := c.Param("id")
	fmt.Println(courseID)

	session := sessions.Default(c)
	userID := session.Get("user-id")
	fmt.Println("...and our user-id is", userID)

	c.HTML(http.StatusOK, "index.html", nil)
}

func CourseHandler(c *gin.Context) {
	fmt.Println("Reached course handler")
	courseID := c.Param("id")
	fmt.Println(courseID)

	session := sessions.Default(c)
	userID := session.Get("user-id")
	fmt.Println("...and our user-id is", userID)

	c.HTML(http.StatusOK, "index.html", nil)
}

/* Package */
func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func RandToken(byteLength int) string {
	b := make([]byte, byteLength)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
