package main
import "fmt"

func main() {
	var jam, menit, i int
	var Tfilm, Twaktu int
	var waktu int

	fmt.Scan(&Tfilm)
	Twaktu = 0
	for i = 1 ; i <= Tfilm ; i++ {
		fmt.Scan(&waktu)
		Twaktu = Twaktu + waktu
	}
		jam = Twaktu / 60
		menit = Twaktu % 60
		
		
		fmt.Print(jam," ", "jam", " ", menit, " ", "menit")
	
}
	