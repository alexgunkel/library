package main

import (
	"fmt"
	"time"
)

type item interface {
	Identifier() (string, error)
	Author() string
}

func printItem(i item, c chan string) {
	identifier, err := i.Identifier()
	if err != nil {
		time.Sleep(time.Second * 3)
		fmt.Println(err)
	}
	fmt.Println(identifier)

	c <- i.Author()
}