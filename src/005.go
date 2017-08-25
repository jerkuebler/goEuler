// It's kind of disgusting how much faster Golang is than Python, considering they don't feel much different to write

package main

import "fmt"

func any(num int) bool {
	for j := 2; j <= 20; j++ {
		if num%j != 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 0
	i := 0

	for num < 1{
		i += 20
		if any(i){
			num = i
		}
	}
	fmt.Println(num)
}
