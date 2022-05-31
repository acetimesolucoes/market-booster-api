package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Enterprise struct {
	ID                 primitive.ObjectID `json:"_id"`
	LegalIdentifier    string             `json:"legalIdentifier"`
	LogoUrl            string             `json:"logoUrl"`
	IsFilial           bool               `json:"isFilial"`
	BusinessName       string             `json:"businessName"`
	FantasyName        string             `json:"fantasyName"`
	RegistrationStatus string             `json:"registrationStatus"`
	LegalNature        string             `json:"legalNature"`
	CNAEPrincipal      string             `json:"CNAEPrincipal"`
	CNAESecondary      string             `json:"CNAESecondary"`
	Users              []string           `json:"users"`
	Customers          []Customer         `json:"customers"`
	Employeers         []Employeer        `json:"employeers"`
}

type EnterpriseInterface interface {
}
