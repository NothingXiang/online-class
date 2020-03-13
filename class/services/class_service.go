package services

import (
	"github.com/NothingXiang/online-class/class"
)

type ClassService interface {
	CreateClass(class *class.Class)
}
