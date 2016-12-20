package main

import (
	"fmt"
	"reflect"
	"time"
)

type job struct {
	Requestor string
	In        int
	Res       int
}

func dbl(name string, inp <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		for {
			select {
			case j := <-inp:
				time.Sleep(time.Second)
				o := job{}
				o.Requestor, o.In = j.Requestor, j.In
				o.Res = 2 * j.In
				out <- o

			}
		}
	}()
	return out
}

//010_OMIT
func main() {
	start := time.Now()
	in := make(chan job, 100)
	numWorkers := 15
	cases := []reflect.SelectCase{}
	for i := 1; i <= numWorkers; i++ {
		out := dbl(fmt.Sprintf("Worker%d", i), in)
		sc := reflect.SelectCase{Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(out)}
		cases = append(cases, sc)
	}

	numReq := 16
	for i := 1; i <= numReq; i++ {
		j := job{Requestor: fmt.Sprintf("Req%d", i), In: i}
		in <- j
	}
	// close(in)

	for i := 1; i <= numReq; i++ {
		_, v, ok := reflect.Select(cases)
		// if i == 0 { // in channel closed
		// 	fmt.Println("Zero selected")
		// 	//break
		// }
		if ok {
			j := v.Interface().(job)
			fmt.Printf("%s: %d => %d\n", j.Requestor, j.In, j.Res)
		} else {
			fmt.Println("Channel Closed")
			// break
		}
	}
	fmt.Printf("%.6f\n", time.Now().Sub(start).Seconds())
}

//020_OMIT
