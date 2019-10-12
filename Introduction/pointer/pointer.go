package pointer

import "fmt"

func Pointer() {
	var yana string = "yana"

	var my *string = &yana

	fmt.Println("yana (value)", yana)
	fmt.Println("yana (address)", &yana)

	fmt.Println("my (value)", my)
	fmt.Println("my (address)", &my)

}
