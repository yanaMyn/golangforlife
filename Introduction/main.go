package main

import (
	"fmt"
)

func main() {

	var x int = 7
	var y string = "tujuh"

	cetak(x, y)

	var b string

	b = indra("ada dana", "ada calon")

	fmt.Printf("hasilnya %s \n", b)

}

func cetak(x int, y string) {
	fmt.Printf("print %d %s \n", x, y)
}

func indra(x string, y string) string {
	var a string

	if x == "ada dana" && y == "ada calon" {
		a = "nikah"
	} else {
		a = "gagal nikah"
	}

	return a
}
