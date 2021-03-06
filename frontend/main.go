package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/NeonSky/studyit/frontend/handlers"
)

var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

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

	authorized := router.Group("/")
	authorized.Use(middleAuthReq())
	{
		authorized.GET("/user/:id", handlers.UserHandler)
		authorized.GET("/course/:id", handlers.CourseHandler)
		authorized.GET("/study/:id", handlers.StudyHandler)
	}

	router.Run(":8080")
}

// Crude middleware
func middleAuthReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user-id")
		if userID == nil {
			handlers.LoginHandler(c)
			c.Abort()
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
