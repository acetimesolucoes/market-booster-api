package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id"`
	FullName string             `json:"full_name"`
	Email    string             `json:"email"`
}

type Users []User
