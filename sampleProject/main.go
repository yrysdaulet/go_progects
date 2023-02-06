package main

import (
	"github.com/yrysdaulet/go_progects/mymodule/mypackage"
	"github.com/yrysdaulet/go_progects/sampleProject/stack"
)

func main() {
	st := stack.Stack{}
	st.Push(5)
	st.Push(6)
	st.Push(7)
	st.Increment()
	st.Print()
	mypackage.PrintHello()
}
