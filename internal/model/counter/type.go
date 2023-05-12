package counter

const (
	CounterRange = 1_000_000
	CounterPath  = "/counter"
)

type (
	Counter struct {
		Start   uint64
		Current uint64
		End     uint64
		Version int32
	}
)
