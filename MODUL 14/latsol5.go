package main

import (
	"fmt"
	"math"
)

func main() {
	var totebag1, totebag2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&totebag1, &totebag2)

		if totebag1 < 0 || totebag2 < 0 || totebag1+totebag2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}
		selisih := math.Abs(totebag1 - totebag2)
		oleng := selisih >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
	}
}
