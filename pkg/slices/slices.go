package slices

import "strconv"

func DeleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
func Contains(arr []string, item string) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}
func GetMiddleValueConvertToInt(arr []string) int {
	m := len(arr) / 2
	i, _ := strconv.Atoi(arr[m])
	return i
}
func GetMiddleValue(arr []int) int {
	m := len(arr) / 2
	return arr[m]
}
