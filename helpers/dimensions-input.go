package helpers

import (
	"fmt"

	"github.com/SamBithrey/coding-software-challenge/validators"
)

func DimensionsInput() (width, height, length float64, windows, doors int) {
	var answer string

	fmt.Println("Please input the dimensions of the room in meters!")
	fmt.Print("Width: ")
	fmt.Scan(&width)
	fmt.Print("Height: ")
	fmt.Scan(&height)
	fmt.Print("Length: ")
	fmt.Scan(&length)
	fmt.Println("Do you have any windows or doors? [y/n]")
	fmt.Scan(&answer)
	if answer == "y" {
		fmt.Print("Windows: ")
		fmt.Scan(&windows)
		fmt.Print("Doors: ")
		fmt.Scan(&doors)
	}

	isValid := validators.ValidateInput(width, height, length, windows, doors)

	if isValid {
		return width, height, length, windows, doors
	}

	fmt.Println("Incorrect dimensions input!")
	return DimensionsInput()
}
