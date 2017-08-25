package main

import "fmt"

func reverse_int(n int) int {
	// Realizing learning Python meant I didn't actually learn the underlying math. Ho-hum.
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func main() {

	max_num := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			new_num := i * j
			if new_num == reverse_int(new_num) && new_num > max_num{
				max_num = new_num
			}
		}
	}
	fmt.Println(max_num)
}
