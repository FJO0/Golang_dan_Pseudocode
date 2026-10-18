package main
import "fmt"

func main () {
	var hasil bool
	var a int64 
	var b int64  
	var c int64   
	
	fmt.Scan(&a, &b, &c)
	hasil = a + b + c >= 2
	fmt.Print(hasil)
}