package main

import (
	"fmt"
	"time"
)

//030_OMIT
func main() {
	timeCh := make(chan time.Time, 100)
	pub(timeCh)
	s1 := sub("Sub1")
	s2 := sub("Sub2")
	publish(timeCh, s1, s2) // HL
	select {}               // wait forever
}

//040_OMIT

//010_OMIT
func pub(ch chan<- time.Time) {
	tkr := time.Tick(time.Second)
	go func() {
		for {
			select {
			case t := <-tkr:
				ch <- t
			}
		}
	}()
}

//020_OMIT
//050_OMIT
func sub(name string) chan<- time.Time {
	ch := make(chan time.Time)
	go func() {
		for {
			t := <-ch
			fmt.Printf("%s: %s\n", name, t.Format("15:04:05.000000"))
		}
	}()
	return ch
}

//060_OMIT
//070_OMIT
func publish(ch <-chan time.Time, subs ...chan<- time.Time) {
	go func() {
		for {
			select {
			case t := <-ch:
				for i := 0; i < len(subs); i++ {
					subs[i] <- t
				}
			}
		}
	}()
}

//080_OMIT
