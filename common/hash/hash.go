package hash

//go:generate mockery --name=Hash
type Hash interface {
	Hash(number uint64) string
}
