package users

import "github.com/globalsign/mgo/bson"

const usersC = "users"

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Email     string        `json:"email"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	AvatarURL string        `json:"avatarUrl"`
}
