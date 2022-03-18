package cmd

import (
	"fmt"
	"bookstore/data"
)

// DeleteBook -  Allows to delete a book from database. The function will
// print out the information about deleted book and returns true if
// the book is deleted successfully. Otherwise, it will print out error message
// and returns false.
func DeleteBook(s []string)bool{
	for _, b := range data.BookDatabase.Items{
		if b.ID == s[2]{
			fmt.Println("--------------- Related Book ----------------")
			b.ShowInfo()
			fmt.Printf("The book whose ID number is %v has been deleted.\n", s[2])
			return true
		}
	}
	fmt.Println("Book ID does not exist in database.")
	return false
	
}