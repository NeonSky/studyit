package courses

import "github.com/globalsign/mgo/bson"

type Comment struct {
	ID      bson.ObjectId `bson:"_id"`
	Author  string        `json:"author"`
	Text    string        `json:"text"`
	Upvotes int           `json:"upvotes"`
}
