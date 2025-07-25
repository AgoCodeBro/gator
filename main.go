package main

import (
	"fmt"
	"github.com/AgoCodeBro/gator/internal/config"
)


func main() {
	fmt.Printf("I love you kylah!\n")
	c, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}

	if err = config.SetUser("Ago"); err != nil {
		fmt.Printf("%v", err)
	}

	c, err = config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("Database: %v\nUser: %v\n", c.DbURL, c.CurrentUserName)
}
