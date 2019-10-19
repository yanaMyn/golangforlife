package main

import (
	//"class-golang/golangforlife/Introduction/adhy"
	//"class-golang/golangforlife/Introduction/cetak"
	"class-golang/golangforlife/Introduction/handler"
	"class-golang/golangforlife/Introduction/model"
	//"fmt"
	//"class-golang/golangforlife/Introduction/indra"
	//"class-golang/golangforlife/Introduction/pointer"
	//"class-golang/golangforlife/Introduction/yana"
	//"fmt"
)

func main() {

	/*
		var x int = 7
		var y string = "tujuh"

		cetak.Cetak(x, y)

		var b string

		b = indra.Indra("ada dana", "ada calon")

		fmt.Printf("hasilnya %s \n", b)

		_, z := adhy.Adhy()

		fmt.Printf("adalah: %s \n", z)

		p := yana.Yana(1, 4, 6)
		fmt.Printf("rata-rata : %f \n", p)

		pointer.Pointer()


	*/

	var createdPerson *model.Person
	createdPerson = handler.Create()

	handler.Edit(createdPerson)

}
