package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rajan170/todo-cli"
)

const (
	todoFile = ".todos.json"
)

func error_handling(err error) error {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		return err
	}
	return nil
}

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("cpl", 0, "marked as completed")
	delete := flag.Int("del", 0, "delete a todo")
	list_all := flag.Bool("ls", false, "lists all todos")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {

	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			error_handling(err)
		}

		todos.Add(task)
		err = todos.Store(todoFile)
		error_handling(err)

	case *complete > 0:
		err := todos.Complete(*complete)
		error_handling(err)
		err = todos.Store(todoFile)
		error_handling(err)

	case *delete > 0:
		err := todos.Delete(*delete)
		error_handling(err)
		err = todos.Store(todoFile)
		error_handling(err)

	case *list_all:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("Text Body Can't be Empty")
	}
	return text, nil
}
