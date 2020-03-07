package models

type Class struct {
	ID   string
	Name string

	School string

	// 班主任
	MasterID string
}
