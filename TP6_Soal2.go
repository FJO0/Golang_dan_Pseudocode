package main
import "fmt"

func main() {
	var W, Tawal, Saving, TSaving, i int

	fmt.Scan(&W, &Tawal, &Saving)
	TSaving = 0
	
	for i = 1 ; i <= W ; i++ {
		TSaving = TSaving + Saving
		Tawal = Tawal + TSaving
	}
	
	fmt.Print(Tawal)
	
}