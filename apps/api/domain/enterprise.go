package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Enterprise struct {
	ID                 primitive.ObjectID `bson:"_id",omitempty json:"_id",omitempty`
	LegalIdentifier    string             `bson:"legalIdentifier" json:"legalIdentifier"`
	LogoUrl            string             `bson:"logoUrl" json:"logoUrl"`
	IsFilial           bool               `bson:"isFilial" json:"isFilial"`
	BusinessName       string             `bson:"businessName" json:"businessName"`
	FantasyName        string             `bson:"fantasyName" json:"fantasyName"`
	RegistrationStatus string             `bson:"registrationStatus" json:"registrationStatus"`
	LegalNature        string             `bson:"legalNature" json:"legalNature"`
	CNAEPrincipal      string             `bson:"CNAEPrincipal" json:"CNAEPrincipal"`
	CNAESecondary      string             `bson:"CNAESecondary" json:"CNAESecondary"`
	Users              []User             `bson:"users" json:"users"`
	Customers          []Customer         `bson:"customers" json:"customers"`
	Employeers         []Employeer        `bson:"employeers" json:"employeers"`
}

type Enterprises []Enterprise
