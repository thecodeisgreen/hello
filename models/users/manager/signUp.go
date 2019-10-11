package manager

import (
	"errors"
	"hello/models/users"

	"go.mongodb.org/mongo-driver/bson"
)

func SignUp(sessionID string, email string, password string) *User {
	user := users.GetOne(bson.M{"email": email})
	if user != nil {
		return errors.New("user already defined")
	}

	ID := users.Create(email, password)
	user := users.GetOneByID(ID)
	user.SetSessionID(sessionID)
	user.Save()
	return user
}
