package main

import "fmt"

func main() {
	min, max := 0, 100
	var input string
	fmt.Println("Please guess a number between 0 - 100 in your mind")

	for min < max {
		mid := (max + min) / 2
		fmt.Printf("Does the number is <= %d\n", mid)
		fmt.Scanf("%s", &input)
		if input != "" && input[0] == 'y' {
			max = mid
		} else {
			min = mid + 1
		}
	}

	fmt.Printf("the number is %d\n", max)

}
