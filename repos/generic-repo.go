package repos

type GenericRepository[T any] interface {
	Create(T) T
	GetList() []T
	GetOne(uint) (T, error)
	Update(uint, T) (T, error)
	DeleteOne(uint) (bool, error)
}
