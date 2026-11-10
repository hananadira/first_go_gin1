package main

import "fmt"

func main() {
	var siswa int
	var gedung string

	fmt.Scanln(&siswa)
	if siswa >= 0 && siswa <= 40 {
		gedung = "TULT"
	} else if siswa > 40 && siswa <= 60 {
		gedung = "Tokong Nanas"
	} else if siswa > 60 && siswa <= 100 {
		gedung = "Cacuk"
	} else {
		gedung = "Kapasitas ruangan tidak mencukupi"
	}

	fmt.Println(gedung)
}