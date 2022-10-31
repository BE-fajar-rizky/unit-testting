package main

import "testing"

func TestPrintLuas(t *testing.T) {
	// your code here
	t.Run("Test fungsi Luas Segitiga", func(t *testing.T) {
		alas := 20
		tinggi := 25
		aktual := PrintLuas(float64(alas), float64(tinggi))
		ekpekstasi := 500 / 2

		if aktual != float64(ekpekstasi) {
			t.Error("Hasil Penambahan tidak sesuai= ", aktual)
		}
	})

}
