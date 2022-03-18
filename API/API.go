package API

import (
	"fmt"
	"bookstore/cmd"
	"strings"
)

// API - Calls the functions which are list, search,
// buy etc. The function will take command line arguments
// as parameter. The first argument (actually it is second but
// the first argument is nothing but program itself.) is command
// which the user will enter and wants to do operations from 
// program. There is ErrorCases function which is clearly explained
// in errorcases.go file. All of the commands are set into map which 
// is api map. With the arguments from users, relative commands will
// be activated. If user doesnt enter a command or makes a mistake
// some errors will be printed out.
func API(s []string){

	if len(s) > 1{
		fmt.Println("api working...")
		ErrorCases(s)

		command := strings.ToLower(s[1])

		apiMap := map[string]func([]string)bool{
			"list": 	cmd.ListBook,
			"search": 	cmd.SearchBook,
			"get":		cmd.GetBook,
			"delete":	cmd.DeleteBook,
			"buy":		cmd.BuyBook,
		}
		if _, ok := apiMap[command]; ok{
			apiMap[command](s)
		}else{
			ErrorCases("Check the command again")
		}
	}else{
		ErrorCases("Please enter a command")
	}
}