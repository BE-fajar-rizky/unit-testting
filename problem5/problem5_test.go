package main

import "testing"

func TestPrintLuasPermukaan(t *testing.T) {
	t.Run("Test fungsi Luas Permukaan Tabung", func(t *testing.T) {
		radius := 4
		tinggi := 20
		aktual := PrintLuasPermukaan(float64(radius), float64(tinggi))
		ekpekstasi := 80

		if aktual != float64(ekpekstasi) {
			t.Error("Hasil Penambahan tidak sesuai= ", aktual)
		}
	})

}
