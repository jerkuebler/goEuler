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

func main(){

	x := 600851475143
	max_prime := 0
	max_num := math.Pow(float64(x), 0.5)
	fmt.Println(max_num)
	for i := 0; i < int(max_num); i++{
		if math.Mod(float64(x), float64(i)) == 0{
			if is_prime(i){
				max_prime = i
			}
		}
	}
	fmt.Println(max_prime)
}