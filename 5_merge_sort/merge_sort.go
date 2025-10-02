/*

1. It uses divide and conquer approach.
2. The array is divided into two equal halves. Each half is recursively processed.
3. At the end off each recursive step, there are two sorted sub-arrays.
4. These two sorted sub-arrays are merged into a sorted one.

*/


package main

import "fmt"

// MergeSort is the main function that initiates the merge sort process.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // Base case: an array with 0 or 1 element is already sorted.
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])  // Recursively sort the left half
	right := MergeSort(arr[mid:]) // Recursively sort the right half

	return merge(left, right) // Merge the sorted halves
}

// merge combines two sorted arrays into a single sorted array.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left slice
	result = append(result, left[i:]...)
	// Append any remaining elements from the right slice
	result = append(result, right[j:]...)

	return result
}

func main() {
	unsortedArray := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", unsortedArray)

	sortedArray := MergeSort(unsortedArray)
	fmt.Println("Sorted array:", sortedArray)
}