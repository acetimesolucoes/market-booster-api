package domain

import "time"

type Customer struct {
	ID              string
	Email           string
	EnterpriseId    string
	FullName        string
	LegalNature     string
	LegalIdentifier string
	Phones          []string
	CreatedDate     time.Time
	CreatedBy       string
	UpdatedDate     time.Time
	UpdatedBy       string
}

type Customers []Customer
