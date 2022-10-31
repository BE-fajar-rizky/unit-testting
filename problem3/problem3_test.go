package main

import "testing"

func TestPrintNama(t *testing.T) {
	// your code here
	t.Run("panggil nama", func(t *testing.T) {

		nama := "kobar"
		aktual := PrintNama(string(nama))
		ekpekstasi := "kobar"

		if aktual != ekpekstasi {
			t.Error("nama tidak sesuai= ", aktual)
		}
	})

}
