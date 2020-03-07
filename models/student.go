package models

type Student struct {
	ID string

	Name string

	Sex sex
}

type sex int

const (
	boy sex = iota
	girl
)
