package main

import (
	"flag"
	"fmt"
	todo "github.com/n-hiraha/cli-todo"
	"os"
)

const (
	todoFile = ".todos.json"
)

func main()  {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del:= flag.Int("del", 0, "delete a todo")
	list:= flag.Bool("list", false, "list todos")

	flag.Parse()
	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del> 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()
	default:
		fmt.Println(os.Stderr, "invalid command")
		os.Exit(0)
	}

}
