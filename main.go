package main

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

// Character sets for coupon generation
// Digits exclude 6 and 9 to avoid confusion with 8
var allowedDigits = "01234578"

// Letters exclude E, F, U, V, I, and O to avoid confusion with numbers or each other
var allowedLetters = "ABCDGHJKLMNPQRSTWXYZ"

var allowedSymbols = "!@#$%"

func main() {
	// Coupon pattern: 'A' for letter, 'D' for digit, 'S' for symbol
	pattern := "AA-DDDD-AA"
	// Number of unique coupons to generate
	couponCount := 10000

	coupons, err := generateUniqueCoupons(pattern, couponCount)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Export generated coupons to a CSV file
	err = exportCouponsToCSV("coupons.csv", coupons)
	if err != nil {
		log.Fatalf("Failed to export coupons to CSV: %v", err)
	}
}

// generateUniqueCoupons creates a list of unique coupon codes based on the pattern and desired amount.
func generateUniqueCoupons(pattern string, amount int) ([]string, error) {
	maxPossible := calculateMaxPossibleCombinations(pattern)
	if amount > maxPossible {
		return nil, fmt.Errorf("can't generate %d unique coupons with pattern '%s'. Max possible: %d", amount, pattern, maxPossible)
	}

	uniqueCoupons := make(map[string]bool)
	var coupons []string

	for len(coupons) < amount {
		code, err := generateCouponFromPattern(pattern)
		if err != nil {
			return nil, err
		}

		if !uniqueCoupons[code] {
			coupons = append(coupons, code)
			uniqueCoupons[code] = true
		}
	}

	return coupons, nil
}

// generateCouponFromPattern creates a single coupon code based on the pattern.
// 'A' = letter, 'D' = digit, 'S' = symbol, any other character is used as-is.
func generateCouponFromPattern(pattern string) (string, error) {
	var builder strings.Builder

	for _, char := range pattern {
		switch char {
		case 'D':
			digit, err := randomCharFromSet(allowedDigits)
			if err != nil {
				return "", err
			}
			builder.WriteByte(digit)
		case 'A':
			letter, err := randomCharFromSet(allowedLetters)
			if err != nil {
				return "", err
			}
			builder.WriteByte(letter)
		case 'S':
			symbol, err := randomCharFromSet(allowedSymbols)
			if err != nil {
				return "", err
			}
			builder.WriteByte(symbol)
		default:
			builder.WriteRune(char) // Use literal character
		}
	}

	return builder.String(), nil
}

// randomCharFromSet returns a random character from the provided set.
func randomCharFromSet(charSet string) (byte, error) {
	max := big.NewInt(int64(len(charSet)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charSet[n.Int64()], nil
}

// calculateMaxPossibleCombinations returns the maximum number of unique codes for a pattern.
func calculateMaxPossibleCombinations(pattern string) int {
	total := 1

	for _, char := range pattern {
		switch char {
		case 'D':
			total *= len(allowedDigits)
		case 'A':
			total *= len(allowedLetters)
		case 'S':
			total *= len(allowedSymbols)
		}
	}

	return total
}

// exportCouponsToCSV writes the list of coupons to a CSV file with one coupon per line.
func exportCouponsToCSV(filename string, coupons []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, coupon := range coupons {
		err := writer.Write([]string{coupon})
		if err != nil {
			return err
		}
	}

	return nil
}
