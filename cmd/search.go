package cmd

import (
	"fmt"
	"strings"
	"bookstore/data"
)

// SearchBook - Allows the user to search a book with full name or
// a word and returns boolean (if the books exists in database, it
// returns true). If the book exists in database, the information about
// the book will be printed out. Furthermore, if there are more than
// one book which contain the searching word, the function will
// print out books which are related to that word.
func SearchBook(l []string)bool{
	fmt.Println("search working...")
	s := stringConcatenator(l)
	fmt.Println("------------- Related Books ----------------")
	fmt.Printf("----------- Searched for(%v) --------------\n", s)
	var count = 0
	for _, b := range data.BookDatabase.Items{
		if strings.Contains(strings.ToLower(b.Name), s){
			b.ShowInfo()
			count += 1
		}
	}
	if count != 0{
		return true
	}else{
		fmt.Println("There is no related book")
		return false
	}
}

// stringComcatenator - Returns concatenated arguments from cli as string.
func stringConcatenator(l []string)string{
	sep := " "
	var s string
	
	for i:=2; i<len(l); i++ {
		if i == len(l) - 1{
			s += l[i]
		}else{
		s += l[i] + sep
		}
	}
	return s
}