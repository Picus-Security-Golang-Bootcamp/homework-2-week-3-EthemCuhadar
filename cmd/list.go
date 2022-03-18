package cmd

import (
	"fmt"
	"bookstore/data"
)

// ListBook - Prints out the books which are stored in
// database and returns boolean (if there is no book 
// in database it returns false).
func ListBook(s []string)bool{
	if len(data.BookDatabase.Items) == 0{
		return false
	}
	fmt.Println("--------------- Books ----------------")
	for _, b := range data.BookDatabase.Items{
		b.ShowInfo()
	}
	return true
}