package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       		primitive.ObjectID 	`json:"_id"`
	FullName 		string             	`json:"full_name"`
	Email    		string             	`json:"email"`
	EnterpriseId 	string 				`json:"enterprise_id"`
	CreatedAt		time.Time 			`json:"created_at"`
	UpdatedAt		time.Time 			`json:"updated_at"`
}

type Users []User
