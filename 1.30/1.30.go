package main

import "fmt"

var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var search int = 6

func main() {
	result, index := BinarySearch(data, search)
	fmt.Println(result, index)
}

func BinarySearch(data []int, search int) (bool, int) {
	size := len(data)
	var mid int
	low := 0
	high := size - 1

	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == search {
			return true, mid
		} else {
			if data[mid] < search {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false, 0
}
