package main

import "fmt"

func main() {
	slice := newSlice()
	printEvenOdd(slice)
}

func newSlice() []int {
	intSlice := []int{}
	for i := 0; i <= 10; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func printEvenOdd(slice []int) {
	for _, item := range slice {
		if item%2 == 0 {
			fmt.Println("even")
		} else if item%2 == 1 {
			fmt.Println("odd")
		}
	}
}
