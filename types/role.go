package types

//go:generate stringer -type=Role
type Role uint8

const (
	Operator Role = iota
	PlatformManager
	Consumer
)
