package repositories

import "acetime-business-crm/apps/api/domain"

type EnterpriseRepository interface {
	Insert(enterprise *domain.Enterprise) (*domain.Enterprise, error)
}
