package utils

// ContainsInSlice check if an element is included within a "slice"
// Takes two parameters that must be of the same type and returns a boolean

func ContainsInSlice[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
