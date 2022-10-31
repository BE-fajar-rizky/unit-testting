package main

import "fmt"

func PrintLuas(alas, tinggi float64) float64 {
	// your code here

	Luas := float64(alas) * float64(tinggi) / 2
	return Luas
}

func main() {
	var alas float64 = 20
	var tinggi float64 = 25

	fmt.Println(PrintLuas(alas, tinggi))
}
