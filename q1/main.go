package main

import (
	"fmt"
	"time"
)

type job struct {
	Name string
	Time time.Time
}

//030_OMIT
func main() {
	p := push("Push1")
	worker("Work1", 1, p) // Fast worker
	worker("Work2", 2, p) // Slower worker
	select {}
}

//040_OMIT

//010_OMIT
func push(name string) <-chan job {
	ch := make(chan job, 1000)
	tkr := time.Tick(time.Second)
	go func() {
		for {
			select {
			case t := <-tkr:
				ch <- job{name, t}
			}
		}
	}()
	return ch
}

//020_OMIT
//050_OMIT
func worker(name string, n int, in <-chan job) {
	go func() {
		for {
			s := <-in
			fmt.Printf("%s: %s  latency:%.6f\n", name,
				s.Time.Format("04:05.000000"), time.Now().Sub(s.Time).Seconds())
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Printf("%s: runtime:%.6f\n", name, time.Now().Sub(s.Time).Seconds())
		}
	}()
}

//060_OMIT
