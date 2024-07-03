package main

import "github.com/oagarian/gosv/app"

func main() {
	err := app.Run()
	if err != nil {
        panic(err)
    }
}