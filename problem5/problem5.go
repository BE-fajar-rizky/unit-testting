package main

import "fmt"

func PrintLuasPermukaan(radius, tinggi float64) float64 {
	// your code here
	Lp := float64(radius) * float64(tinggi)
	return Lp
}

func main() {
	var T float64 = 20
	var r float64 = 4

	fmt.Println(PrintLuasPermukaan(r, T))
}
