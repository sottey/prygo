package main

import "github.com/sottey/prygo/pry"

func a(b int) {
	c := 5
	pry.Pry()
	_ = c
}

func main() {
	a(5)
}
