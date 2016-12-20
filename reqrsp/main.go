package main

import (
	"fmt"
	"reflect"
	"time"
)

func dbl(i int) int {
	time.Sleep(time.Second)
	return 2 * i
}
func dblg(i int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		ch <- 2 * i
	}()
	return ch
}

func seq() {
	start := time.Now()
	for i := 1; i <= 5; i++ {
		fmt.Println(dbl(i))
	}
	fmt.Printf("proc time: %.6f\n", time.Now().Sub(start).Seconds())
}

func goRtn() {
	start := time.Now()
	cases := []reflect.SelectCase{}
	for i := 1; i <= 5; i++ {
		sc := reflect.SelectCase{Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(dblg(i))}
		cases = append(cases, sc)
	}
	for i := 1; i <= 5; i++ {
		j, v, ok := reflect.Select(cases)
		if ok {
			fmt.Printf("dblg %d = %d\n", j+1, v.Int())
		}
	}
	fmt.Printf("proc time: %.6f\n", time.Now().Sub(start).Seconds())
}

//010_OMIT
func main() {
	seq()
	goRtn()
}

//020_OMIT
