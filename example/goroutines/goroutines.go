package main

import (
	"fmt"
	"time"

	"github.com/sottey/prygo/pry"
)

func prying() {
	fmt.Println("PRYING!")
}

func main() {
	c := make(chan bool)
	go func() {
		prying()
		pry.Pry()
		c <- true
	}()
	<-c
	for {
		time.Sleep(time.Second)
	}
}
