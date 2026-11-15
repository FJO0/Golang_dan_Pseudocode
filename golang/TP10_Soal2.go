package main
import "fmt"

func main () {
	var status string
	var t, h, k int64
	
	fmt.Scan(&t, &h, &k)
	
	if t*75 <= h*100 && k >= 8 {
		status = "BOLEH UAS"
	} else {
		status = "TIDAK BOLEH UAS"
    }
	
	fmt.Print(status)
}