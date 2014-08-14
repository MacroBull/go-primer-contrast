package main

import (
	"fmt"
	"time"
)

const LIM = 9999999

func main() {

	tStart := time.Now()

	primes := [LIM + 1]int{2, 3}
	index := 2

	var curIsPrime bool

	//for n := 5; n < LIM; n++ {
	for n := 5; n < LIM; n += 2 {
		//for _, m := range primes {
		for i := 0; ; i++ {
			m := primes[i]
			if 0 == n%m {
				curIsPrime = false
				break
			} else if m*m > n {
				curIsPrime = true
				break
			}
		}

		if curIsPrime {
			primes[index] = n
			index++
		}
		//fmt.Printf("%v\n", primes)
	}

	fmt.Printf("%s:\t%d \n", time.Since(tStart), index)

}
