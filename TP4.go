package main

import "fmt"

func main() {
	var angka int64
	var hasil bool

	fmt.Scan(&angka)
	hasil = angka >= 1500
	fmt.Print(hasil)
}