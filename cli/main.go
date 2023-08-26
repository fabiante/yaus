package main

import (
	"fmt"
	"github.com/fabiante/yaus/app"
	"os"
)

func main() {
	args := os.Args[1:]
	input := args[0]

	service := app.NewService()

	url, err := service.ShortenURL(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(url)
}
