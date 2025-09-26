package utils

import "strconv"

// mustAtoi is a helper function to convert a string to an int, panicking on error.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}