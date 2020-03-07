package models

const (
	Chinese Subject = iota
	Maths
	English
	Science
	Physics
	Chemistry
	Biology
	Politics
	Moral
	History
	Geography
	Natural
	Sports
	Technology
	Art
	Music
	Other
)

type Teacher struct {
	// 唯一id
	ID string

	// 教授科目
	Subject Subject
}

type Subject int
