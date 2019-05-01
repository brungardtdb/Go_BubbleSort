package main

import (
	"fmt"
)

// Initialize slice
var mySlice = make([]int, 8)

func main() {

	// Add values to slice
	mySlice[0] = 10
	mySlice[1] = 60
	mySlice[2] = 120
	mySlice[3] = 2
	mySlice[4] = 15
	mySlice[5] = 85
	mySlice[6] = 100
	mySlice[7] = 620

	// Print slice values
	fmt.Println(mySlice)

	// Sort slice with bubble sort algorithm
	newSlice := bubbleSort(mySlice)

	// Print sorted slice values
	fmt.Println(newSlice)
}

// Method for bubble sort algorithm
func bubbleSort(param []int) []int{

	// Pass slice to local variable
	thisSlice := param

	// Sentinal value for loop
	var Sentinal = true

	// Sort slice
	for isSorted := true; isSorted; isSorted = Sentinal {
	// Set sentinal value to true at the beginning of each iteration
	Sentinal = false

		// Loop through items in slice
		for i:=0; i < (len(thisSlice) - 1); i++ {

			// If second item is smaller than first item,
			// Swap items in slice
			if(thisSlice[i + 1] < thisSlice[i]) {

				// Swap items
				temp1 := thisSlice[i + 1]
				temp2 := thisSlice[i]
				thisSlice[i + 1] = temp2
				thisSlice[i] = temp1

				// Set Sentinal value to true if a swap is necessary
				Sentinal = true
			}

		}
	}

	return thisSlice
}
