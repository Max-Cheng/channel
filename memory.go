package channel

type memoryChannel[T any] struct {
	c chan T
}

func (m *memoryChannel[T]) Write(element T) error {
	m.c <- element
	return nil
}

func (m *memoryChannel[T]) Read() (T, error) {
	return <-m.c, nil
}

func (m *memoryChannel[T]) Close() error {
	close(m.c)
	return nil
}

func newMemoryChannel[T any](capacity int64) (Channel[T], error) {
	return &memoryChannel[T]{
		c: make(chan T, capacity),
	}, nil
}
