package main

import (
	"fmt"

	"github.com/sottey/prygo/pry"
)

func main() {
	for i := 0; i < 10; i++ {
		pry.Pry()
	}
	fmt.Println("DUCK")
}
