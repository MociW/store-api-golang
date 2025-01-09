package util

import (
	"fmt"
	"strings"
	"unicode"
)

func generateID(product string) string {
	var filteredProduct []string

	for _, char := range product {
		if unicode.IsLetter(char) {
			filteredProduct = append(filteredProduct, string(char))
		}
	}

	result := strings.ToUpper(strings.Join(filteredProduct[0:3], ""))

	return result
}

func generateSizeID(size string) string {
	size = strings.ToUpper(size)
	var sizeID string

	// Map common clothing sizes to 3-letter abbreviations
	switch size {
	case "S", "SMALL":
		sizeID = "SMA"
	case "M", "MEDIUM":
		sizeID = "MED"
	case "L", "LARGE":
		sizeID = "LAR"
	case "XL":
		sizeID = "XLA"
	case "XXL":
		sizeID = "XXL"
	default:
		if isNumeric(size) {
			sizeID = fmt.Sprintf("%03s", size)
		} else {
			sizeID = fmt.Sprintf("%-3.3s", size)
		}
	}

	return sizeID
}

func isNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func SKUGenerator(userID string, product string, size string, color string) string {

	productID := generateID(product)
	sizeID := generateSizeID(size)
	colorID := generateID(color)
	result := fmt.Sprintf("%s-%s-%s-%s", productID, sizeID, colorID, userID[:8])

	return result
}
