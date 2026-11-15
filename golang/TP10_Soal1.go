package main
import "fmt"

func main () {
	var jarak, harga int64 
	
	fmt.Scan(&jarak)
	
	if jarak <= 5 {
		harga = 10000
	} else if jarak > 5 {
		harga = 14000
	}
	
	fmt.Print(harga)
}