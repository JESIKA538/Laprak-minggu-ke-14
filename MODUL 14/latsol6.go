package main

import "fmt"

func main() {
	var K float64
	fmt.Scan(&K)

	fK := (4*K + 2) * (4*K + 2)/ ((4*K + 1) * (4*K + 3))
	fmt.Printf("Nilai f(K) = %.10f\n",  fK)

}