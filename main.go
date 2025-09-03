package main

import "fmt"

func main() {
	A := []int{1, 3, 5}
	B := []int{3, 2, 6}
	Out := []int{}

	for _, num := range A {
		if Contains(num, B) {Out = append(Out, num)}
	}

	fmt.Println(Out)
}

func Contains(number int, numbers []int) bool {
	// Или можно использовать модуль slices
	for _, num := range numbers {
		if num == number {
			return true
		}
	}
	return false
}
