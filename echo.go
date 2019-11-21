package main

import (
	"fmt"
	"os"
)

func main()  {
	var s = ""
	for i, agv := range os.Args[1:] {
		s += " " + agv
		fmt.Println(i, agv)
	}
	println(s)
}
