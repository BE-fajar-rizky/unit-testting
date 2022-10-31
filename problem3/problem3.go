package main

import "fmt"

func PrintNama(nama string) string {
	// your code here
	panggilan := nama
	return panggilan
}

func main() {
	var nama string = "kobar"
	fmt.Println(PrintNama(nama))
}
