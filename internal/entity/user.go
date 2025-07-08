package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
}

func NewUser(email, pasword string) *User {
	return &User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: pasword,
	}
}
