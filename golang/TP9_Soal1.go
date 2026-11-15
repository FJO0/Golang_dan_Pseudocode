package main
import "fmt"

func main () {
	var N int64
	var gedung string
	
	fmt.Scan(&N)
	if N > 0 &&  N <= 40 {
		gedung = "TULT"
	}
	if N >= 41 && N <= 60 {
		gedung = "Tokong Nanas"
	}
	if N >= 61 && N <= 100 {
		gedung = "Cacuk"
	}
	if N > 100 {
		gedung = "Kapasitas ruangan tidak mencukupi"
	}
	fmt.Println(gedung)
}	