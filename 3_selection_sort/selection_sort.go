/*

1. Iterate over the array to find the smallest element.
2. At the end of each cycle, this smallest element is swapped with the first one.
3. The next cycle excludes the last positioned smallest item.

*/

package main

import "fmt"

func selectionSort(inputArray []int) {
	for i := 0; i < len(inputArray)-1; i++ {
		min := i
		for j := i + 1; j < len(inputArray); j++ {
			if inputArray[j] < inputArray[min] {
				min = j
			}
		}
		if min != i {
			temp := inputArray[min]
			inputArray[min] = inputArray[i]
			inputArray[i] = temp
		}
		fmt.Println("Sorted Array:", inputArray)
	}

	// fmt.Println("sorted array:", inputArray)
}

func main() {
	var numberOfElements, temporaryVariable int
	fmt.Println("Enter the value for number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var arr = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element:\n", i)
		fmt.Scanf("%d\n", &temporaryVariable)
		arr[i] = temporaryVariable
	}
	fmt.Println("arr:", arr)
	selectionSort(arr)
}
