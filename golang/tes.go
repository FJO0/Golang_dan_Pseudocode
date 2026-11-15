package main
import "fmt"

func main () {
	var a, b, c, d, determinan float64
	
	fmt.Scan(&a, &b, &c, &d)
	determinan = (a * d) - (b*c)
	
	fmt.Print(determinan)
}
	