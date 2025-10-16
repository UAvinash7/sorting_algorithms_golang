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

func main() {
	data := []int{12, 34, 54, 2, 3}
	fmt.Println("Unsorted array:", data)
	ShellSort(data)
	fmt.Println("Sorted array:", data) // Output: Sorted array: [2 3 12 34 54]

	data2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Unsorted array:", data2)
	ShellSort(data2)
	fmt.Println("Sorted array:", data2) // Output: Sorted array: [1 2 3 4 5 6 7 8 9]
}
