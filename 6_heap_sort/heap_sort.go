/*

1. It keeps the items in a min-heap data structure.
2. It swaps the root with the right most child node, removes it from the heap, and insert it in the final array.
3. Then it moves the smallest remaining item to the root.
4. It stops when the root is empty.

*/

package main

import "fmt"

func linearSort(inputArray []int, dataElement *int) {
	var i int
	//var elementFound int = 0
	for i = 0; i < len(inputArray); i++ {
		if inputArray[i] == *dataElement {
			fmt.Printf("Element found at index: %d and its position is: %d and its value is: %d\n", i, i+1, inputArray[i])
			//elementFound = 1
			break
		}
	}
	if i == len(inputArray) {
		fmt.Println("Element not found")
	}
}

func main() {
	var numberOfElements, temp, data int
	fmt.Println("Enter the value of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var arr = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element:\n", i)
		fmt.Scanf("%d\n", &temp)
		arr[i] = temp
	}
	fmt.Println("array:", arr)
	fmt.Println("Enter the value of element that need to be searched in the input array:")
	fmt.Scanf("%d\n", &data)
	linearSort(arr, &data)

}
