package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/studyit/frontend/handlers"
)

var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	router := gin.Default()

	// Cookie store
	store := sessions.NewCookieStore([]byte(handlers.RandToken(64)))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	router.Use(sessions.Sessions("goquestsession", store))
	router.LoadHTMLGlob("public/*.html")

	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)

	router.StaticFile("", "./public/index.html")
	router.StaticFile("index.html", "./public/index.html")
	router.StaticFile("bundle.js", "./public/bundle.js")
	router.StaticFile("bundle.css", "./public/bundle.css")

	authorized := router.Group("/study")
	authorized.Use(middleAuthReq())
	{
		authorized.GET("/:id", handlers.StudyHandler)
	}

	router.Run(":8080")
}

// Crude middleware
func middleAuthReq() gin.HandlerFunc {
	fmt.Println("MIDDLEWARE")
	return func(c *gin.Context) {
		fmt.Println("Checking user id")
		session := sessions.Default(c)
		valid := session.Get("user-id")
		if valid == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
		c.Next()
	}
}
