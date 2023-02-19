package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/daicang/gi/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
