package data

import (
	"homework-2-week-3-EthemCuhadar/books"
)

// Decleare books variables.
var Book001 = books.Book{}
var Book002 = books.Book{}
var Book003 = books.Book{}
var Book004 = books.Book{}
var Book005 = books.Book{}

// Decleare book list variable.
var BookDatabase = books.BookList{}

func init(){
	
	// Set necessary fields of books.
	Book001.SetID("1001")
	Book001.SetName("Don Quixote")
	Book001.SetPageNumber(1072)
	Book001.SetStockNumber(10)
	Book001.SetPrice(11.69)
	Book001.SetStockCode("543891")
	Book001.SetISBN("9142437239")
	Book001.SetAuthor("Miguel De Cervantes Saavedra")
	
	Book002.SetID("1002")
	Book002.SetName("One Hundred Years of Solitude")
	Book002.SetPageNumber(417)
	Book002.SetStockNumber(15)
	Book002.SetPrice(14.50)
	Book002.SetStockCode("394536")
	Book002.SetISBN("2060883286")
	Book002.SetAuthor("Gabriel Garcia Marquez")
	
	Book003.SetID("1003")
	Book003.SetName("War and Peace")
	Book003.SetPageNumber(1296)
	Book003.SetStockNumber(12)
	Book003.SetPrice(15.69)
	Book003.SetStockCode("264159")
	Book003.SetISBN("1400079985")
	Book003.SetAuthor("Leo Tolstoy")
	
	Book004.SetID("1004")
	Book004.SetName("Nineteen eighty four")
	Book004.SetPageNumber(328)
	Book004.SetStockNumber(13)
	Book004.SetPrice(7.48)
	Book004.SetStockCode("660315")
	Book004.SetISBN("9780451524935")
	Book004.SetAuthor("George Orwell")
	
	Book005.SetID("1005")
	Book005.SetName("Madame Bovary")
	Book005.SetPageNumber(384)
	Book005.SetStockNumber(9)
	Book005.SetPrice(14.49)
	Book005.SetStockCode("940031")
	Book005.SetISBN("6451418506")
	Book005.SetAuthor("Gustave Flaubert")
	
	// Store the books in book list.
	BookDatabase.Items = append(BookDatabase.Items, Book001, Book002, Book003, Book004, Book005)
}