package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"packer/box"
)

//go:generate go run gen.go

func main() {

	b := box.GetDefault()
	c, err := b.Get(".env")
	if err != nil {
		panic(err)
	}

	env, err := godotenv.Unmarshal(string(c))
	if err != nil {
		panic(err)
	}

	fmt.Println("env mod:", env["ENV_MODE"])
}
