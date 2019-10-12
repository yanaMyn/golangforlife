package yana

func Yana(n ...int) float32 {
	var total int

	for _, x := range n {
		total += x
	}

	p := float32(total / len(n))
	return p
}
