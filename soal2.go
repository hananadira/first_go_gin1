package main

import "fmt"

func main() {
	var N int
	var bil string

	fmt.Scanln(&N)
	if N % 2 == 0 {
		bil = "Genap"
	} else {
		bil = "Ganjil"
	}
	fmt.Println(bil)
}