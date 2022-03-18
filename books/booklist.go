package books

import (
	"sync"
)

type BookList struct{
	Items 	[]Book
	mux		sync.RWMutex
}