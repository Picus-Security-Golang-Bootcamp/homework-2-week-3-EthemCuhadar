package books

import (
	"sync"
)

type BookList struct{
	Items 	[]Book
	Mux		sync.RWMutex
}