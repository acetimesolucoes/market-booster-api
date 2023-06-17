package repository

type Repository[T any] interface {
	FindAll(page int64, limit int64) ([]*T, error)
	FindOneById(id string) (T, error)
	Save(entity T) error
	Update(id string, entity T) error
	Delete(id string) error
}
