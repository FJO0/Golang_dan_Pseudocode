package main
import "fmt"

func main (){
	var N int64
	var hasil string
	
	fmt.Scan(&N)
	
	if N % 2 == 1 {
		hasil = "ganjil"
	}
	if N % 2 == 0 {
		hasil = "genap"
	}
	
	fmt.Print(hasil)

}
		
