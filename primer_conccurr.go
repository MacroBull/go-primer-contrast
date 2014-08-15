package main

import (
	"fmt"
	"runtime"
	"time"
)

const LIM = 9999999

var primes []int // nonono primes are not sorted, giving wrong result
var v [LIM + 1]bool
var c, h chan int
var na int = 3

func checkPrime(n int) {
	for _, m := range primes {
		if m*m > n {
			c <- -n
			return
		} else if 0 == n%m {
			c <- n
			return
		}
	}
}

func prepare() {
	var r int
	if r = <-c; r < 0 {
		r = -r
		primes = append(primes, r)
	}
	v[r] = true
	for v[na] {
		na = na + 2
	}
}

func main() {
	PR := 4 //runtime.NumCPU()
	runtime.GOMAXPROCS(PR)
	tStart := time.Now()

	primes = []int{3, 5}
	v[3] = true
	v[5] = true
	c = make(chan int, PR)

	for i := 0; i < PR; i++ {
		c <- 0
	}

	////for n := 5; n < LIM; n++ {
	for n := 7; n < LIM; n += 2 {
		prepare()
		for na*na <= n {
			prepare()
		}
		go checkPrime(n)
	}

	for na < LIM {
		prepare()
	}

	//fmt.Printf("\n%v\n", primes)
	fmt.Printf("%s:\t%d \n", time.Since(tStart), len(primes)+1)

}
