/*

Shell Sort is a generalized version of Insertion Sort algorithm.
It first sorts elements that are far apart from each other and successfully reduces the interval between the elements to be sorted.

*/

package main

// ShellSort sorts an array of integers using the Shell Sort algorithm.
func ShellSort(arr []int) {
	n := len(arr)

	// Start with a large gap, then reduce the gap
	for gap := n / 2; gap > 0; gap /= 2 {
		// Do a gapped insertion sort for this gap size.
		// The first gap elements arr[0..gap-1] are already in gapped order
		// keep adding one more element until the entire array is gap sorted
		for i := gap; i < n; i++ {
			// add arr[i] to the elements that have been gap sorted
			// save arr[i] in temp and make a hole at position i
			temp := arr[i]

			// shift earlier gap-sorted elements up until the correct location for arr[i] is found
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			// put temp (the original arr[i]) in its correct location
			arr[j] = temp
		}
	}
}
