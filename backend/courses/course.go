package courses

import "github.com/globalsign/mgo/bson"

const coursesC = "courses"

type Course struct {
	ID          bson.ObjectId `bson:"_id"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Level       string        `json:"level"`
	Posts       []Post        `json:"posts"`
}
