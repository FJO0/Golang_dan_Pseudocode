package main
import "fmt"

func main() {
	var s, pitaKeliling, pitaTotal int

	// Menerima masukan panjang sisi
	fmt.Scan(&s)

	// Menghitung keliling dan total pita
	pitaKeliling = 4 * s
	pitaTotal = pitaKeliling + 20

	// Menampilkan output
	fmt.Println(pitaKeliling)
	fmt.Println(pitaTotal)

	// Penjelasan
	fmt.Println("Penjelasan :")
	fmt.Println("Pita keliling = ", pitaKeliling)
	fmt.Println("Pita total = ", pitaTotal)
}

