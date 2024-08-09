package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirusxcoder/todo-app"
)

const (
	todoFile = ".todos.json"
)

func main() {
	// Corrected: flag.Bool(name:"add", value: false, usage:"add a new todo") -> flag.Bool("add", false, "add a new todo")
	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample Todo")     // Removed incorrect syntax: task:
		err := todos.Store(todoFile) // Declare err properly here
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		// Corrected: fmt.Fprintln(os.Stdout, a...:"invalid command") -> fmt.Fprintln(os.Stdout, "invalid command")
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
