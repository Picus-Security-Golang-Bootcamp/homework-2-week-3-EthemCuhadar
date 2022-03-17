package books

// Book - Struct for books with related fields.
type Book struct {
	ID			int
	Name		string
	PageNumber	int
	StockNumber	int
	Price		float64
	StockCode	int
	ISBN		int
	Author		string
}

// Decleare books
var Book001 = Book{}
var Book002 = Book{}
var Book003 = Book{}
var Book004 = Book{}
var Book005 = Book{}

// SetID - Sets the ID field and returns the book.
func (b *Book) SetID(i int) *Book {
	b.ID = i
	return b
}

// SetName - Sets the Name field and returns the book.
func (b *Book) SetName(s string) *Book {
	b.Name = s
	return b
}

// SetPageNumber - Sets the PageNumber field and returns the book.
func (b *Book) SetPageNumber(i int) *Book {
	b.PageNumber = i
	return b
}

// SetStockNumber - Sets the StockNumber field and returns the book.
func (b *Book) SetStockNumber(i int) *Book {
	b.StockNumber = i
	return b
}

// SetPrice - Sets the Price field and returns the book.
func (b *Book) SetPrice(f float64) *Book {
	b.Price = f
	return b
}

// SetStockCode - Sets the StockCode field and returns the book.
func (b *Book) SetStockCode(i int) *Book {
	b.StockCode = i
	return b
}

// SetISBN - Sets the ISBN field and returns the book.
func (b *Book) SetISBN(i int) *Book {
	b.ISBN = i
	return b
}

// SetAuthor - Sets the Author field and returns the book.
func (b *Book) SetAuthor(s string) *Book {
	b.Author = s
	return b
}