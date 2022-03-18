# Homework 2 | Week 3

Published story about concurrency concept can be seen in link below.

[Medium Concurrency Link](https://medium.com/@iethemcuhadar/all-you-need-to-know-about-concurreny-in-golang-9486bc149a8f)

![Golang Image](golang.png)

-------------------------------------------------------------------

This is a very simple CLI (command line interface) application in Go programming language. 

The program has 5 commands:

```
list

search <bookName>

get <bookID>

delete <bookID>

buy <bookID> <quantity>
```

--------------------------------------------------------

## Usage of commands

* Usage of **list** command can be seen below.

```[console]
$ go run main.go list
```

```[echo]
--------------- Books ----------------
Book ID:      1001
Name:         Don Quixote
Page Number:  1072
Stock Number: 10
Price:        11.69 $
Stock Code:   543891
ISBN:         9142437239
Author:       Miguel De Cervantes Saavedra
--------------------------------------------
Book ID:      1002
Name:         One Hundred Years of Solitude
Page Number:  417
Stock Number: 15
Price:        14.5 $
Stock Code:   394536
ISBN:         2060883286
Author:       Gabriel Garcia Marquez
--------------------------------------------
Book ID:      1003
Name:         War and Peace
Page Number:  1296
Stock Number: 12
Price:        15.69 $
Stock Code:   264159
ISBN:         1400079985
Author:       Leo Tolstoy
--------------------------------------------
Book ID:      1004
Name:         Nineteen eighty four
Page Number:  328
Stock Number: 13
Price:        7.48 $
Stock Code:   660315
ISBN:         9780451524935
Author:       George Orwell
--------------------------------------------
Book ID:      1005
Name:         Madame Bovary
Page Number:  384
Stock Number: 9
Price:        14.49 $
Stock Code:   940031
ISBN:         6451418506
Author:       Gustave Flaubert
--------------------------------------------
```

* Furthermore, the usage of **search bookname** can seen below. The user can enter a word or even a character to search a book.

```[console]
$ go run main.go search solitude
```

```[console]
------------- Related Books ----------------
----------- Searched for(solitude) --------------
Book ID:      1002
Name:         One Hundred Years of Solitude
Page Number:  417
Stock Number: 15
Price:        14.5 $
Stock Code:   394536
ISBN:         2060883286
Author:       Gabriel Garcia Marquez
--------------------------------------------
```

* The usage of **get bookID** is shown below.

```[console]
go run main.go get 1001
```

```[console]
--------------- Related Book ----------------
Book ID:      1001
Name:         Don Quixote
Page Number:  1072
Stock Number: 10
Price:        11.69 $
Stock Code:   543891
ISBN:         9142437239
Author:       Miguel De Cervantes Saavedra
--------------------------------------------
```

* The usage of **delete bookID** can be seen below.

```[console]
go run main.go delete 1001
```

```[console]
--------------- Related Book ----------------
Book ID:      1001
Name:         Don Quixote
Page Number:  1072
Stock Number: 10
Price:        11.69 $
Stock Code:   543891
ISBN:         9142437239
Author:       Miguel De Cervantes Saavedra
--------------------------------------------
The book whose ID number is 1001 has been deleted.
```

* The usage of **buy bookID quantity** is shown below.

```[console]
go run main.go buy 1002 5
```

```[console]
The purchase process is done successfully
Book ID:      1002
Name:         One Hundred Years of Solitude
Page Number:  417
Stock Number: 10
Price:        14.5 $
Stock Code:   394536
ISBN:         2060883286
Author:       Gabriel Garcia Marquez
```

-------------------------------------------

## License

MIT