package main

var data = []int{2, 5, 4}

var search = 5
var search_not_present = 3

func main() {
	SequentialSearch(data, search)
	SequentialSearch(data, search_not_present)
}

func SequentialSearch(data []int, search int) (bool, int) {

	for index, value := range data {
		if search == value {
			return true, index
		}
	}
	return false, 0
}
