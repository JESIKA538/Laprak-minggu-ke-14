package main

import "fmt"

func main() {
	var percobaan int
	fmt.Print("Jumlah percobaan: ")
	fmt.Scan(&percobaan)

	for i := 1; i <= percobaan; i++ {
		var warna1, warna2, warna3, warna4 string

		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		fmt.Printf("%s\n%s\n%s\n%s\n", warna1, warna2, warna3, warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			fmt.Println("BERHASIL: false")
			return
		}
	}

	fmt.Println("BERHASIL: true")
}
