package main

import (
	"fmt"

	"github.com/beevik/guid"
)

func main() {
	fmt.Println(guid.New().String())
}
