package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
