package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	result := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			result++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil\n", result)

}