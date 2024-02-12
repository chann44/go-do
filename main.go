package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
	id    int
	title string
}

func main() {
	fmt.Println("Go DO - A CLI PROGRAM TO MANAGE YOUR TODOS")
	fmt.Println("Version : 0.0.1")
	fmt.Println("COMMANDS: ")
	fmt.Println("help = 'h' ")
	fmt.Println("list all todos = 'la' ")
	fmt.Println("clear all todos = 'c' ")
	fmt.Println("create todo = 'a' ")
	fmt.Println("delete todo = 'd {id}' ")
	fmt.Println("quit = ':q'")

	var cmd string = ""
	todo := make([]Todo, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for cmd != ":q" {
		fmt.Scan(&cmd)
		if cmd == "a" {
			println("what is your item called?:")
			scanner.Scan()
			item := scanner.Text()
			todo = append(todo, Todo{
				id:    1,
				title: item,
			})
			fmt.Printf("your Todos: %+v\n", todo)
		}
	}

}
