package types

//go:generate stringer -type=State
type State uint8

const (
	Frozen State = iota
	Active
)
