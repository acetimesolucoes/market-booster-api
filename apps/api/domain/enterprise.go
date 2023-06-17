package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Enterprise struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BusinessName       string             `bson:"business_name" json:"business_name"`
	CNAE               []string           `bson:"CNAE" json:"CNAE"`
	CreatedAt          time.Time          `bson:"created_at" json:"created_at"`
	Customers          Customers          `bson:"customers" json:"customers"`
	Employeers         Employeers         `bson:"employeers" json:"employeers"`
	FantasyName        string             `bson:"fantasy_name" json:"fantasy_name"`
	IsFilial           bool               `bson:"is_filial" json:"is_filial"`
	LegalIdentifier    string             `bson:"legal_identifier" json:"legal_identifier"`
	LegalNature        string             `bson:"legal_nature" json:"legal_nature"`
	LogoUrl            string             `bson:"logo_url" json:"logo_url"`
	RegistrationStatus string             `bson:"registration_status" json:"registration_status"`
	UsersId            []string           `bson:"users_id" json:"users_id"`
	UpdatedAt          time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewEnterprise(business_name string, cnae []string, fantasy_name string) *Enterprise {
	return &Enterprise{
		ID:           primitive.NewObjectID(),
		BusinessName: business_name,
		CNAE:         cnae,
		CreatedAt:    time.Now(),
		FantasyName:  fantasy_name,
	}
}

// type Enterprises []*Enterprise
