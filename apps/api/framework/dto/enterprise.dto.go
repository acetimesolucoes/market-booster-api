package dto

type EnterpriseCreateDTO struct {
	BusinessName       string   `bson:"business_name" json:"business_name"`
	CNAE               []string `bson:"CNAE" json:"CNAE"`
	FantasyName        string   `bson:"fantasy_name" json:"fantasy_name"`
	IsFilial           bool     `bson:"is_filial" json:"is_filial"`
	LegalIdentifier    string   `bson:"legal_identifier" json:"legal_identifier"`
	LegalNature        string   `bson:"legal_nature" json:"legal_nature"`
	LogoUrl            string   `bson:"logo_url" json:"logo_url"`
	RegistrationStatus string   `bson:"registration_status" json:"registration_status"`
}

type EnterpriseUpdateDTO struct {
	BusinessName       string   `bson:"business_name" json:"business_name"`
	CNAE               []string `bson:"CNAE" json:"CNAE"`
	FantasyName        string   `bson:"fantasy_name" json:"fantasy_name"`
	IsFilial           bool     `bson:"is_filial" json:"is_filial"`
	LegalIdentifier    string   `bson:"legal_identifier" json:"legal_identifier"`
	LegalNature        string   `bson:"legal_nature" json:"legal_nature"`
	LogoUrl            string   `bson:"logo_url" json:"logo_url"`
	RegistrationStatus string   `bson:"registration_status" json:"registration_status"`
}
