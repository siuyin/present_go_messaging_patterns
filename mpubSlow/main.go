package main

import (
	"fmt"
	"time"
)

//030_OMIT
func main() {
	p1 := pub("Pub1")
	s1 := sub("Sub1")

	for {
		select {
		case t := <-p1:
			s1 <- t
		}
	}
}

//040_OMIT

//010_OMIT
func pub(name string) <-chan string {
	ch := make(chan string) // make be buffered // HL
	//012_OMIT
	tkr := time.Tick(100 * time.Millisecond)
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
	ch := make(chan string) // make me buffered // HL
	//052_OMIT
	go func() {
		for {
			fmt.Printf("%s: %s\n", name, <-ch)
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}

//060_OMIT
