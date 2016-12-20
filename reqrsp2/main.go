package main

import (
	"fmt"
	"time"
)

//050_OMIT
type job struct {
	Requestor chan<- int // HL
	In        int
}

//060_OMIT

//030_OMIT
func dbl(jc <-chan job) {
	go func() {
		for {
			j := <-jc
			time.Sleep(time.Second) // slow worker
			j.Requestor <- j.In * 2 // reply channel // HL
		}
	}()
}

//040_OMIT
//070_OMIT
func request(jc chan<- job, i int, done chan<- struct{}) {
	go func() {
		start := time.Now()
		rply := make(chan int)
		j := job{Requestor: rply, In: i}
		jc <- j
		fmt.Printf("2*%02d = %02d s:%s run:%.6f\n", i, <-rply, // Get reply // HL
			start.Format("15:04:05.000000"),
			time.Now().Sub(start).Seconds())
		done <- struct{}{}
	}()
}

//080_OMIT

//010_OMIT
func main() {
	dblQ := make(chan job, 100)
	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		dbl(dblQ)
	}
	//012_OMIT
	//013_OMIT
	done := make(chan struct{})
	numReq := 10
	for i := 1; i <= numReq; i++ {
		request(dblQ, i, done)
	}
	//014_OMIT
	start := time.Now()
	n := 0
MAIN:
	for {
		select {
		case <-done:
			n++
			if n >= numReq {
				close(dblQ)
				fmt.Printf("Exiting. Completed %d jobs\n", n)
				break MAIN
			}
		}
	}
	fmt.Printf("total run time: %.6f\n", time.Now().Sub(start).Seconds())
}

//020_OMIT
