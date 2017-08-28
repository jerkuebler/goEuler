package main

import "fmt"

func main() {

	result := 0
	new_num := 0
	temp1 := 0
	temp2 := 1
	for temp1 < 4000000 {
		new_num = temp1 + temp2
		temp2 = temp1
		temp1 = new_num
		if new_num%2 == 0 {
			result += new_num
		}
	}
	fmt.Println(result)
}
