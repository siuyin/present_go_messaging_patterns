package main

import (
	"fmt"
	"time"
)

//030_OMIT
func main() {
	p := pub()
	s := sub()
	for {
		select {
		case t := <-p:
			s <- t
		}
	}
}

//040_OMIT

//010_OMIT
func pub() <-chan string {
	ch := make(chan string)
	tkr := time.Tick(time.Second)
	go func() {
		for {
			select {
			case t := <-tkr:
				ch <- t.String()
			}
		}
	}()
	return ch
}

//020_OMIT
//050_OMIT
func sub() chan<- string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	return ch
}

//060_OMIT
