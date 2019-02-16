package courses

import "github.com/globalsign/mgo/bson"

type Post struct {
	ID       bson.ObjectId `bson:"_id"`
	Author   string        `json:"author"`
	Title    string        `json:"title"`
	Text     string        `json:"text"`
	Upvotes  int           `json:"upvotes"`
	Comments []Comment     `json:"comments"`
}
