package main

import (
	"os"

	tokenizer "github.com/sujalcodes3/stylesheet-tokenizer/tokenizer"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	t := tokenizer.New(string(dat))
	e := t.Tokenize()
	check(e)

	t.PrintTokens()
}
