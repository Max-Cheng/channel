package channel

import "errors"

type Type int

const (
	TypeFile Type = iota
	TypeMemory
)

var (
	ErrInitCapacityError = errors.New("init capacity error")
)
