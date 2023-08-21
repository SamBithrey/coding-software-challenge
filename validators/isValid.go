package validators

func ValidateInput(width, height, length float64, windows, doors int) bool {
	isValidWidth := width > 0
	isValidHeight := height > 0
	isValidLength := length > 0
	isValidWindow := windows >= 0
	isValidDoors := doors >= 0

	if isValidDoors && isValidHeight && isValidLength && isValidWidth && isValidWindow {
		return true
	} else {
		return false
	}
}
