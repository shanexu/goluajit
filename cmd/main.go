package main

import (
	"github.com/shanexu/goluajit"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Errorf("main file.lua\n")
		os.Exit(1)
	}
	file := os.Args[1]
	_, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	s := goluajit.NewState()
	s.LoadFile(file)
	fmt.Println(s.PCall(0, 0, 0))
	s.GetGlobal("address")
	s.ToString(-1)
}
