package counter

type Repository interface {
	GetCurrentCounter() (uint64, error)
}
