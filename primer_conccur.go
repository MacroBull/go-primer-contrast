package main

import (
	"fmt"
	"runtime"
	"time"
)

const LIM = 9999999

var primes []int
var c chan int

//var currents [4]int // no lock for valid checking

func checkPrime(n int) {
	for _, m := range primes {
		if m*m > n {
			primes = append(primes, n)
			<-c
			return
		} else if 0 == n%m {
			<-c
			return
		}
	}
}

func main() {
	PR := runtime.NumCPU() * 2
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(PR)
	tStart := time.Now()

	primes = []int{3, 5}
	//c = make(chan int)
	c = make(chan int, PR)

	////for n := 5; n < LIM; n++ {
	for n := 7; n < LIM; n += 2 {
		c <- 1
		go checkPrime(n)
		//print("c", n, "\t")
	}

	for i := 0; i < PR; i++ {
		c <- 1
	}

	//fmt.Printf("%v", primes)
	fmt.Printf("%s:\t%d \n", time.Since(tStart), len(primes)+1)

}
