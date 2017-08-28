// Also spits out a slice that contains all of the factors of x, because I needed to learn slices

package main

import (
	"fmt"
	"math"
)

func is_prime(x int) bool {

	max_num := math.Pow(float64(x), 0.5)
	for i := 2; i < int(max_num); i++ {
		if math.Mod(float64(x), float64(i)) == 0 {
			return false
		}
	}
	return true
}

func main() {

	x := 600851475143
	factor_list := []int{}
	max_prime := 0
	max_num := math.Pow(float64(x), 0.5)
	for i := 0; i < int(max_num); i++ {
		if math.Mod(float64(x), float64(i)) == 0 {
			factor_list = append(factor_list, i)
			if is_prime(i) {
				max_prime = i
			}
		}
	}
	fmt.Println(max_prime)
	fmt.Println(factor_list)
}
