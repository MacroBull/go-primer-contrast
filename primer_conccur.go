package main

import (
	"fmt"
	"runtime"
	"time"
)

const LIM = 9999999

var primes []int
var c chan bool

//var currents [4]int // no lock for valid checking

func checkPrime(n int) {
	for _, m := range primes {
		if m*m > n {
			break
		} else if 0 == n%m {
			c <- false
			return
		}
	}
	primes = append(primes, n)
	c <- true
	return
}

func main() {
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	tStart := time.Now()

	primes = []int{2, 3}
	c = make(chan bool, 4)

	//for n := 5; n < LIM; n++ {
	for n := 5; n < LIM; n += 2 {
		go checkPrime(n)
		<-c
	}

	fmt.Printf("%s:\t%d \n", time.Since(tStart), len(primes))

}
