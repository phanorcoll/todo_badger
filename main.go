package main

import (
	"fmt"
	"github.com/phanorcoll/todo_badger/config"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(config.EnvVariables.PORT)
}
