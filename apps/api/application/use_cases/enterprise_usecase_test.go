package enterprise_use_cases_test

import (
	"testing"
	"time"

	enterpriseUseCase "acetime.com.br/business-crm/apps/api/application/use_cases"
	"acetime.com.br/business-crm/apps/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var enterpriseId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	enterpriseId = oid.Hex()

	enterprise := domain.Enterprise{
		ID:                 oid,
		LegalIdentifier:    "00000000000000000",
		LogoUrl:            "https://teste.com.br/imagem.png",
		IsFilial:           true,
		BusinessName:       "Empresa Teste LTDA",
		FantasyName:        "Empresa Teste",
		RegistrationStatus: "ACTIVE",
		LegalNature:        "PJ",
		CNAE:               []string{"0000", "0001", "0002"},
		UsersId:            []string{},
		Customers:          []domain.Customer{},
		Employeers:         []domain.Employeer{},
	}

	err := enterpriseUseCase.Create(enterprise)

	if err != nil {
		t.Error("Teste de persistencia de dados da empresa falhou!")
		t.Fail()
	} else {
		t.Log("Teste de criação de empresa finalizado com sucesso!")
	}
}

func TestFindById(t *testing.T) {

	_, err := enterpriseUseCase.FindOneById(enterpriseId)

	if err != nil {
		t.Error("Teste de busca por id retornou erro")
		t.Fail()
	} else {
		t.Log("Teste de busca por id finalizado com sucesso!")
	}
}

func TestFindAll(t *testing.T) {

	enterprises, err := enterpriseUseCase.FindAll()

	if err != nil {
		t.Error("Teste de busca de empresas falhou!")
		t.Fail()
	}

	if len(enterprises) == 0 {
		t.Error("Quantidade de empresas igual a zero.")
		t.Fail()
	} else {
		t.Log("Teste de busca de empresas finalizado com sucesso!")
	}
}

func TestUpdate(t *testing.T) {

	enterprise := domain.Enterprise{
		BusinessName: "",
		CNAE:         []string{},
		UpdatedAt:    time.Now(),
	}

	err := enterpriseUseCase.Update(enterpriseId, enterprise)

	if err != nil {
		t.Error("Erro ao atualizar empresa")
		t.Fail()
	} else {
		t.Log("Teste de atualização finalizou com sucesso!")
	}
}

func TestDelete(t *testing.T) {

	err := enterpriseUseCase.Delete(enterpriseId)

	if err != nil {
		t.Error("Erro ao deletar empresa")
		t.Fail()
	} else {
		t.Log("Teste de deleção da empresa finalizado com sucesso!")
	}
}
