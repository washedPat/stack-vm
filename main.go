package main

import (
	"flag"
	"github.com/washedPat/stack-vm/vm"
)

func main() {
	path := flag.String("path", "", "Path to the file to run")
	flag.Parse()
	program, err := vm.Parse(*path)
	if err != nil {
		panic(err)
	}
	stack := vm.NewStack()

	stack.Run(program)
}
