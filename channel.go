package channel

type Channel[T any] interface {
	Write(element T) (err error)
	Read() (T, error)
	Close() error
}

func NewChannel[T any](t Type, capacity int64, elementType T) (Channel, error) {
	switch t {
	case TypeFile:
		return newFileChannel(capacity)
	case TypeMemory:
		return newMemoryChannel[T](elementType, capacity)
	default:
		return nil, ErrInitCapacityError
	}
}
