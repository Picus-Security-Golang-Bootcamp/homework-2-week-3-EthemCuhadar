package API

import (
	"fmt"
	"strings"
)

// ErrorCases - Prints out error message when user dont enter
// a valid statement on command line arguments. The errors are considered
// for all commands seperately.
func ErrorCases(s []string)bool{

	command := strings.ToLower(s[1])
	
	if command == "list"{
		if len(s) > 2{
			Error("Check the usage of list again")
			return true
		}
	}else if command == "search"{
		if len(s) < 3{
			Error("Enter the book name to search")
			return true
		}
	}else if command == "get"{
		if len(s) < 3{
			Error("Enter the book ID go get information")
			return true
		}else if len(s) > 3{
			Error("Please enter book ID only")
			return true
		}
	}else if command == "delete"{
		if len(s) < 3{
			Error("Enter the book ID to delete")
			return true
		}
	}else if command == "buy"{
		if len(s) != 4{
			Error("Check the book ID and quantity again")
			return true
		}
	}
	return false
}

// Error - prints out usage of program when user make a 
// mistake.
func Error(s string){
	fmt.Println(s)
	fmt.Println(err)
}

// Usage of commands
var err = `usage: 
go run main.go list
go run main.go get <bookID>
go run main.go search <bookName>
go run main.go delete <bookID>
go run main.go buy <bookID> <quantity>`