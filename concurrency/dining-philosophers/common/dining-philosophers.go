package common

import (
	"sync"
)

type Fork struct {
	Id   int
	Lock *sync.Mutex
}

func NewFork(id int) *Fork {
	lock := new(sync.Mutex)
	return &Fork{id, lock}
}
