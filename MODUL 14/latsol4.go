package main

import "fmt"

func main() {
	var pita, result string
	var bunga int

	for i := 1; ; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&result)

		if result == "SELESAI" {
			break
		}
		if bunga > 0 {
			pita += " - "
		}
		pita += result
		bunga++
	}
	if bunga > 0 {
		pita+= " - "
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bunga)
}
