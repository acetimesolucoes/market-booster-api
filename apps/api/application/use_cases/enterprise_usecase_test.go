package enterprise_use_cases_test

import (
	"testing"

	enterpriseUseCase "acetime.com.br/business-crm/apps/api/application/use_cases"
	"acetime.com.br/business-crm/apps/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {

	enterprise := domain.Enterprise{
		ID:                 primitive.NewObjectID(),
		LegalIdentifier:    "",
		LogoUrl:            "",
		IsFilial:           true,
		BusinessName:       "",
		FantasyName:        "",
		RegistrationStatus: "",
		LegalNature:        "",
		CNAEPrincipal:      "",
		CNAESecondary:      "",
	}

	err := enterpriseUseCase.Create(enterprise)

	if err != nil {
		t.Error("")
		t.Fail()
	} else {
		t.Log("")
	}

}

func TestFindAll(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
