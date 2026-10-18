package main

import "fmt"

func main() {
	var n int64
	var hari int64
	fmt.Scan(&n)
	hari = ((4+ n - 1) % 7) + 1
	fmt.Print(hari)
}