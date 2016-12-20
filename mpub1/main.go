package main

import (
	"fmt"
	"time"
)

//030_OMIT
func main() {
	p1 := pub("Pub1")
	p2 := pub("Pub2")
	s1 := sub("Sub1")
	s2 := sub("Sub2")

	for {
		select {
		case t := <-p1:
			s1 <- t
			s2 <- t
		case t := <-p2:
			s1 <- t
			s2 <- t
		}
	}
}

//040_OMIT

//010_OMIT
func pub(name string) <-chan string {
	ch := make(chan string)
	tkr := time.Tick(time.Second)
	go func() {
		for {
			select {
			case t := <-tkr:
				ch <- fmt.Sprintf("%s: %s", name, t.String())
			}
		}
	}()
	return ch
}

//020_OMIT
//050_OMIT
func sub(name string) chan<- string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Printf("%s: %s\n", name, <-ch)
		}
	}()
	return ch
}

//060_OMIT
