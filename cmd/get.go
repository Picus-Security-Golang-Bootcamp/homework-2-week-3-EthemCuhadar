package cmd

import (
	"fmt"
	"homework-2-week-3-EthemCuhadar/data"
)

// GetBook - Allows the user to see information about a book by entering
// the book ID. The function returns boolean. If the ID which are typed by
// user doesnt match in database, the function will return false. Otherwise,
// the user can see all information (relative fields) about the book.
func GetBook(s []string)bool{
	data.BookDatabase.Mux.Lock()
	defer data.BookDatabase.Mux.Unlock()
	for _, b := range data.BookDatabase.Items{
		if b.ID == s[2]{
			fmt.Println("--------------- Related Book ----------------")
			b.ShowInfo()
			return true
		}
	}
	fmt.Println("Book ID does not exist in database.")
	return false
}