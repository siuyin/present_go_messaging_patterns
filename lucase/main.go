package main

import (
	"fmt"
	"strings"
	"time"
)

// Lower Upper Case
// input: AbC, output: abcABC

//010_OMIT
func lCase() (chan string, chan string) {
	sc := make(chan string, 100)
	ch := make(chan string, 100)
	go func() {
		for {
			s := <-sc
			time.Sleep(time.Second)
			ch <- strings.ToLower(s)
		}
	}()
	return sc, ch
}

//020_OMIT

func uCase() (chan string, chan string) {
	sc := make(chan string, 100)
	ch := make(chan string, 100)
	go func() {
		for {
			s := <-sc
			time.Sleep(time.Second)
			ch <- strings.ToUpper(s)
		}
	}()
	return sc, ch
}

func tCase() (chan string, chan string) {
	sc := make(chan string, 100)
	ch := make(chan string, 100)
	go func() {
		for {
			s := <-sc
			time.Sleep(time.Second)
			o := ""
			switch {
			case len(s) == 0:
				o = ""
			case len(s) == 1:
				o = strings.ToUpper(s[0:1])
			case len(s) > 1:
				o = strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
			}
			ch <- o
		}
	}()
	return sc, ch
}

//090_OMIT
func cons(lc, uc, tc <-chan string) {
	var l, u, t string
	var lSet, uSet, tSet bool
	go func() {
		for {
			select {
			case l = <-lc:
				lSet = true
			case u = <-uc:
				uSet = true
			case t = <-tc:
				tSet = true
			}
			if lSet && uSet && tSet {
				fmt.Printf("Output: %s|%s|%s  %.6f\n", l, u, t, time.Now().Sub(start).Seconds())
				lSet, uSet, tSet = false, false, false
			}
		}
	}()
}

//100_OMIT
func publish(in <-chan string, s1, s2, s3 chan<- string) {
	go func() {
		for {
			s := <-in
			s1 <- s
			s2 <- s
			s3 <- s
		}
	}()
}

var start time.Time

func main() {
	//050_OMIT
	wordQ := make(chan string, 100)
	//060_OMIT
	//070_OMIT
	lIn, lOut := lCase()
	uIn, uOut := uCase()
	tIn, tOut := tCase()
	publish(wordQ, lIn, uIn, tIn)
	cons(lOut, uOut, tOut)
	//080_OMIT

	s := ""
	//030_OMIT
	for {
		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		fmt.Printf("You entered %s\n", s)
		wordQ <- s // HL
		start = time.Now()
	}
	//040_OMIT
}
