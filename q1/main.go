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
	jobQ := make(chan job, 100)
	push("Push1", jobQ)
	//032_OMIT
	worker("Work1", 1, jobQ) // Fast worker
	worker("Work2", 2, jobQ) // Slower worker
	select {}
}

//040_OMIT
func main2() { // HL
	jobQ := make(chan job, 100)
	push("Push1", jobQ)
	//032_OMIT
	for i := 1; i <= 10; i++ { // HL
		name := fmt.Sprintf("Work%d", i)
		speed := 1
		worker(name, speed, jobQ)
	}
	select {}
}

//046_OMIT
//010_OMIT
func push(name string, jc chan<- job) {
	tkr := time.Tick(time.Second)
	go func() {
		for {
			select {
			case t := <-tkr:
				jc <- job{name, t} // HL
			}
		}
	}()
}

//020_OMIT
//050_OMIT
func worker(name string, n int, jc <-chan job) {
	go func() {
		for {
			s := <-jc
			fmt.Printf("%s: %s  latency:%.6f\n", name,
				s.Time.Format("04:05.000000"), time.Now().Sub(s.Time).Seconds())
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Printf("%s: runtime:%.6f\n", name, time.Now().Sub(s.Time).Seconds()) //	 // HL
		}
	}()
}

//060_OMIT
