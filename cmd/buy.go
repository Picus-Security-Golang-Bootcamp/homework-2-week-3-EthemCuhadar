package cmd

import (
	"fmt"
	"homework-2-week-3-EthemCuhadar/data"
	"strconv"
)

// BuyBook - Allows the user to buy a book. The user should enter
// both book ID and quantity. The function will return true if purchase
// is done successfully. Otherwise, it will return false. The quantity entry
// should be a valid number. Moreover, if the quantity is greater than stock
// number of book, a message will be seen to user.
func BuyBook(s []string)bool{
	data.BookDatabase.Mux.Lock()
	defer data.BookDatabase.Mux.Unlock()
	quantity, err := strconv.Atoi(s[3])
	if err != nil{
		fmt.Println("Please enter a valid number")
		return false
	}else{
		for _, b := range data.BookDatabase.Items{
		
			if b.ID == s[2]{
				// stocknumber < quantity
				if b.StockNumber < quantity{
					fmt.Printf("There are only %v books in stocks\n", b.StockNumber)
					return false
				}else if b.StockNumber >= quantity{	// stocknumber >= quantity
					fmt.Println("The purchase process is done successfully")
					b.StockNumber -= quantity
					b.ShowInfo()
					return true
				}
			}
		}
		fmt.Println("Book ID does not exist in database.")
		return false
	}
}