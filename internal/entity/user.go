package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
}

func NewUser(email, pasword, username string) *User {
	return &User{
		ID:       primitive.NewObjectID(),
		UserName: username,
		Email:    email,
		Password: pasword,
	}
}
