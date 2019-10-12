package model

type Person struct {
	ID      int
	Name    string
	Age     int
	Address string
	Dream   Dream
}

type Dream struct {
	Type        string
	Description string
}
