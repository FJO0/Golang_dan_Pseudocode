package main

import "fmt"

func main() {
	var N int64
	var h1, h2, h3 int64
	var sisa int64
	fmt.Scan(&N)
	h1 = N / 3600
	sisa = N % 3600
	h2 = sisa / 60
	h3 = sisa % 60
	
	fmt.Println("jam :", h1, "menit :", h2, "detik :", h3)
}