package courses

import (
	"fmt"
	"github.com/globalsign/mgo/bson"

	"github.com/NeonSky/studyit/backend/util"
)

func RequestCourse(code string) Course {
	c, session := util.FetchCollection(coursesC)
	defer session.Close()

	result := Course{}
	if err := c.Find(bson.M{"code": code}).One(&result); err != nil {
		fmt.Println(err)
	}
	return result
}

func requestCreateCourse(course Course) {
	c, session := util.FetchCollection(coursesC)
	defer session.Close()
	course.ID = bson.NewObjectId()
	c.Insert(&course)
}
