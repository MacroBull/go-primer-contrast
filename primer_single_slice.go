package main

import (
	"fmt"
	"time"
)

const LIM = 9999999

func main() {

	//var primes [LIM + 1]bool

	//primes[2] = true
	//primes[3] = true

	tStart := time.Now()

	primes := []int{2, 3}
	var curIsPrime bool

	//for n := 5; n < LIM; n++ {
	for n := 5; n < LIM; n += 2 {
		for _, m := range primes {
			if m*m > n {
				curIsPrime = true
				break
			} else if 0 == n%m {
				curIsPrime = false
				break
			}
		}
		if curIsPrime {
			primes = append(primes, n)
		}
	}

	fmt.Printf("%s:\t%d \n", time.Since(tStart), len(primes))

}
