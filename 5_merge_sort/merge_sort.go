/*

1. It uses divide and conquer approach.
2. The array is divided into two equal halves. Each half is recursively processed.
3. At the end off each recursive step, there are two sorted sub-arrays.
4. These two sorted sub-arrays are merged into a sorted one.

*/

package main

import "fmt"

func linearSearch(inputArray[]int, element *int) {
	// var i int
	var found int = 0
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == *element {
			fmt.Printf("Element found at index: %d and its position is: %d\n", i, i+1)
			found = 1
			break
		}
	}
	if found == 0 {
		fmt.Println("Element not found in the given array")
	}
}

func main() {
	var numberOfElements, temporaryVariable, searchElement int
	fmt.Println("Enter the value of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var arr = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element\n", i)
		fmt.Scanf("%d\n", &temporaryVariable)
		arr[i] = temporaryVariable
	}
	fmt.Println("arr", arr)
	fmt.Println("Enter the value of element that is to be searched:")
	fmt.Scanf("%d\n", &searchElement)
	linearSearch(arr, &searchElement)
}
