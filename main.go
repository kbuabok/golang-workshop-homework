package main

import (
	"demo/cmd/controller"
	"os"
	"strconv"
	"strings"
)

func main() {
	op := os.Args[1]
	switch op {
	case "add":
		N := strings.Split(os.Args[2], "=")[1]
		a := strings.Split(os.Args[3], "=")[1]
		Age,e := strconv.Atoi(a)
		if e != nil {
			Age = 0
		}
		t := controller.Task{FileName: "./data.json"}
		t.AddList(N,Age)
	case "list":
		t := controller.Task{FileName: "./data.json"}
		t.ShowList()
	}
}
