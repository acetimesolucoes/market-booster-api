package usecases

type UseCase[T any] interface {
	NewUseCase() T
	FindAll(page int64, limit int64) ([]*T, error)
	FindOneById(id string) (T, error)
	Create(entity T) error
	Update(id string, entity T) error
	Delete(id string) error
}
