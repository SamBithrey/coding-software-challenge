package main

import (
	"fmt"

	"github.com/SamBithrey/coding-software-challenge/helpers"
)

var (
	width   float64
	height  float64
	length  float64
	windows int
	doors   int

	floorArea  float64
	roomVolume float64

	maximumPaint float64
	wallPaint    float64
)

func main() {
	fmt.Println("Welcome to my Borwell Software Challenge")
	fmt.Println("######################")

	width, height, length, windows, doors = helpers.DimensionsInput()

	fmt.Println("######################")

	floorArea = helpers.CalculateArea(width, length)
	roomVolume = helpers.CalculateVolume(height, width, length)

	fmt.Printf("Your floor area is %v㎡ \n", floorArea)
	fmt.Printf("Your room volume is %v㎥\n", roomVolume)

	fmt.Println("######################")

	maximumPaint, wallPaint = helpers.CalculatePaint(width, height, length, windows, doors)

	fmt.Println("You will need this much paint to cover the walls.")
	fmt.Printf("Maximum ammount of paint: %v gallons", maximumPaint)
	fmt.Printf("\nIf you take into account the windows and doors: %v gallons", wallPaint)
	fmt.Println("\nPlease bare in mind that this number is based off average window and door size and plan accordingly")
	fmt.Println("######################")
}
