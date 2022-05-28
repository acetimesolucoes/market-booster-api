package repositories

import "acetime.com.br/business-crm/apps/api/domain"

type EnterpriseRepository interface {
	Insert(enterprise *domain.Enterprise) (*domain.Enterprise, error)
}
