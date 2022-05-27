package domain

type Enterprise struct {
	ID                 string      `json:"_id"`
	LegalIdentifier    string      `json:"legalIdentifier"`
	LogoUrl            string      `json:"LogoUrl"`
	IsFilial           string      `json:"isFilial"`
	BusinessName       string      `json:"businessName"`
	FantasyName        string      `json:"fantasyName"`
	RegistrationStatus string      `json:"registrationStatus"`
	LegalNature        string      `json:"legalNature"`
	CNAEPrincipal      string      `json:"CNAEPrincipal"`
	CNAESecondary      string      `json:"CNAESecondary"`
	Users              []string    `json:"users"`
	Customers          []Customer  `json:"customers"`
	Employeers         []Employeer `json:"employeers"`
}
