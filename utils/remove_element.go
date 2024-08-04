package utils

//RemoveElement removes an element within the slice
//Takes two parameters that must be of the same type and returns a new slice

func RemoveElement[T comparable](slice []T, element T) []T {
	index := -1
	for i, v := range slice {
		if v == element {
			index = i
			break
		}
	}
	if index != -1 {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
