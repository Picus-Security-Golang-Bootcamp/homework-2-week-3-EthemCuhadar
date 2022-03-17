package main

import (
	"fmt"
)

func main(){
	Args := os.Args
	fmt.Println(Args[1:])
}