package main

import (
	"fmt"
	"github.com/yrysdaulet/go_progects/mymodule/mypackage"
)

func main() {
	fmt.Println("Hello, Modules!")
	mypackage.PrintHello()

	for true {

		fmt.Println("Hello, Modules!")
	}
}
