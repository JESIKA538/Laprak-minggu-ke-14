package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x > 1 {
		prima := true
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				prima = false
				break
			}
		}
		if prima {
			fmt.Println("prima")
		} else {
			fmt.Println("bukan prima")
		}
	}
}
