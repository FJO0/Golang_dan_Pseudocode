package main
import "fmt"

type data struct {
	nama string
	nim string
	kuis float64
	tugas float64
	praktikum float64
	UAS float64
	UTS float64
	akhir float64
}

func main () {
	var d data
	fmt.Scan(&d.nama, &d.nim, &d.kuis, &d.tugas, &d.praktikum, &d.UAS, &d.UTS)
	d.akhir = 0.15 * d.kuis + 0.15 * d.tugas + 0.20 * d.praktikum + 0.25 * d.UAS + 0.25 * d.UTS
	fmt.Println("NIM :", d.nim) 
	fmt.Println("Nama :", d.nama) 
	fmt.Println("Nilai kuis :", d.kuis) 
	fmt.Println("Nilai tugas :", d.tugas) 
	fmt.Println("Nilai praktikum :", d.praktikum) 
	fmt.Println("Nilai UAS :", d.UAS) 
	fmt.Println("Nilai UTS :", d.UTS) 
	fmt.Println("Nilai akhir :", d.akhir)
}