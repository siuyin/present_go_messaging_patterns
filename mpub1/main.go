package main

import (
	"fmt"
	"time"
)

//030_OMIT
func main() {
	timeCh := make(chan time.Time, 100)
	pub("UTC", timeCh)   // HL
	pub("Local", timeCh) // HL
	s1 := sub("Sub1")
	s2 := sub("Sub2")
	publish(timeCh, s1, s2)
	select {} // wait forever
}

//040_OMIT

//010_OMIT
func pub(name string, ch chan<- time.Time) { // HL
	tkr := time.Tick(time.Second)
	go func() {
		for {
			select {
			case t := <-tkr:
				switch name {
				case "UTC":
					ch <- t.UTC()
				case "Local":
					ch <- t.Local()
				}
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
			fmt.Printf("%s: %s\n", name, t.Format("MST 15:04:05.000000"))
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
