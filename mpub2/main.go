package main

import (
	"fmt"
	"time"
)

//030_OMIT
func main() {
	pubs := fanIn(pub("Pub1"), pub("Pub2"))
	subs := fanOut(sub("Sub1"), sub("Sub2"))
	for {
		select {
		case t := <-pubs:
			subs <- t
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
	ch := make(chan string, 1000) // receive buffer for each subscriber // HL
	go func() {
		for {
			fmt.Printf("%s: %s\n", name, <-ch)
		}
	}()
	return ch
}

//060_OMIT
//070_OMIT
func fanIn(s1, s2 <-chan string) <-chan string {
	out := make(chan string, 1000) // send buffer // HL
	go func() {
		for {
			select {
			case t := <-s1:
				out <- t
			case t := <-s2:
				out <- t
			}
		}
	}()
	return out
}

//080_OMIT
//090_OMIT
func fanOut(o1, o2 chan<- string) chan<- string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case t := <-ch:
				o1 <- t // see sub on how we overcome
				o2 <- t // slow consumers
			}
		}
	}()
	return ch
}

//100_OMIT
