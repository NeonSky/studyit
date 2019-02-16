package users

import (
	"fmt"
	"github.com/globalsign/mgo/bson"

	"github.com/NeonSky/studyit/backend/util"
)

func RequestUser(email string) User {
	c, session := util.FetchCollection(usersC)
	defer session.Close()

	result := User{}
	if err := c.Find(bson.M{"email": email}).One(&result); err != nil {
		fmt.Println(err)
	}
	return result
}

func requestUsers() []User {
	c, session := util.FetchCollection(usersC)
	defer session.Close()

	result := []User{}
	if err := c.Find(nil).All(&result); err != nil {
		fmt.Println(err)
	}
	return result
}

func requestCreateUser(newUser User) {
	c, session := util.FetchCollection(usersC)
	defer session.Close()
	newUser.ID = bson.NewObjectId()
	c.Insert(&newUser)
}

func requestUpdateUser(prevUser User, newUser User) {
	c, session := util.FetchCollection(usersC)
	defer session.Close()
	updateQuery := bson.M{
		"$set": bson.M{
			"email":     newUser.Email,
			"firstName": newUser.FirstName,
			"lastName":  newUser.LastName,
			"avatarUrl": newUser.AvatarURL,
		},
	}
	c.Update(prevUser, updateQuery)
}

func requestDeleteUser(user User) bool {
	c, session := util.FetchCollection(usersC)
	defer session.Close()
	err := c.Remove(&user)
	return err == nil
}
