package main

import "fmt"

func main() {
	array := [9][9]int{}
	fmt.Println(array)
	changeArray(&array)
	fmt.Println(array)
}

func changeArray(arr *[9][9]int) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			arr[i][j] += i + 1
		}
	}
}
