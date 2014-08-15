package main

import (
	"fmt"
	"time"
)

const LIM = 9999999

func main() {

	tStart := time.Now()

	primes := []int{2, 3}

	//for n := 5; n < LIM; n++ {
	for n := 5; n < LIM; n += 2 {
		for _, m := range primes {
			if m*m > n {
				primes = append(primes, n)
				break
			} else if 0 == n%m {
				break
			}
		}
	}

	fmt.Printf("%s:\t%d \n", time.Since(tStart), len(primes))

}
