package courses

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* Public */
func SetupRoutes(router *gin.Engine) {
	router.GET("/courses/:code", getCourseByCode)
	router.GET("/courses/:code/posts", getCoursePosts)
	router.POST("/courses/", createCourse)
}

/* GET */
func getCourseByCode(c *gin.Context) {
	courseCode := c.Param("code")
	course := RequestCourse(courseCode)
	c.JSON(http.StatusOK, course)
}

func getCoursePosts(c *gin.Context) {
	courseCode := c.Param("code")
	posts := RequestCourse(courseCode).Posts
	c.JSON(http.StatusOK, posts)
}

/* POST */
func createCourse(c *gin.Context) {
	var newCourse Course
	json.NewDecoder(c.Request.Body).Decode(&newCourse)
	fmt.Println(newCourse)
	requestCreateCourse(newCourse)
	c.Status(http.StatusOK)
}
