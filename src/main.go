package main

import (
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.lang")
	source := string(bytes)

	tokens

}
