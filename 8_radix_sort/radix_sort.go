/*

1. It finds the maximum value in the array, determining how many digits it has.
2. It iterates over each element in the array considering first unit's place, then ten's place, then hundred's place etc... etc...
3. It arranges the elements based on sorted numbers in that digit's place.

*/

package main

import "fmt"

// getMax finds the maximum value in an array.
func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// countSort sorts the array based on the digit represented by 'exp'.
func countSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10) // For digits 0-9

	// Store count of occurrences in count[]
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}

	// Change count[i] so that count[i] now contains actual position of this digit in output[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	// Copy the output array to arr[], so that arr[] now contains sorted numbers according to current digit
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// RadixSort sorts an array of integers using Radix Sort.
func RadixSort(arr []int) {
	// Find the maximum number to know number of digits
	max := getMax(arr)

	// Do counting sort for every digit. Note that 'exp' is passed
	// as 10^i where i is current digit number
	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
}

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("Original array:", arr)
	RadixSort(arr)
	fmt.Println("Sorted array:", arr)

	arr2 := []int{10, 1, 5, 100, 12}
	fmt.Println("Original array:", arr2)
	RadixSort(arr2)
	fmt.Println("Sorted array:", arr2)
}
