/*

1. It iterates through the array to find the number of elements.
2. It creates a counter array of size (number of elements + 1) with all zero's.
3. It iterates the array and updates the counter array accordingly.
4. It sorts the array using the counter array.

*/

package main

// CountingSort sorts an array of non-negative integers using the counting sort algorithm.
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// 1. Find the maximum element to determine the range of values.
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	// 2. Create a count array to store the frequency of each element.
	// The size of the count array is maxVal + 1 to accommodate values from 0 to maxVal.
	count := make([]int, maxVal+1)
	for _, val := range arr {
		count[val]++
	}

	// 3. Modify the count array to store the cumulative count.
	// This helps in determining the correct sorted position of each element.
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// 4. Create an output array to store the sorted elements.
	output := make([]int, len(arr))

	// 5. Iterate through the input array in reverse to ensure stability.
	// Place each element in its correct sorted position in the output array.
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		// The position in the output array is (cumulative count of val) - 1.
		output[count[val]-1] = val
		// Decrement the count for the current value, as it has been placed.
		count[val]--
	}

	return output
}