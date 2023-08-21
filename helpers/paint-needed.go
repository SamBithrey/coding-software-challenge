package helpers

import "math"

var (
	paintCoverage = 30   // Average amount of wall size that 1 gallon of paint will cover (in meters)
	doorSize      = 1.44 // Average UK door size.
	windowSize    = 0.5  // Average UK window size (based off 1m height and 0.5m width)
)

func CalculatePaint(width, height, length float64, windows, doors int) (maximumPaint, wallPaint float64) {
	totalArea := width*height*2 + length*height*2

	maximumPaint = math.Ceil(totalArea*100/float64(paintCoverage)) / 100 // Calculates gallons of paint to 2 sigfigs

	doorArea := doorSize * float64(doors)
	windowArea := windowSize * float64(windows)

	wallArea := totalArea - doorArea - windowArea

	wallPaint = math.Ceil(wallArea*100/float64(paintCoverage)) / 100 // Calculates gallons of paint to 2 sigfigs

	return maximumPaint, wallPaint
}
