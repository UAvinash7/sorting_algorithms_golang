/*

1. In each cycle an item is compared with all other elements preceeding it, in reverse order.
2. If the element is smaller than the compared element, this compared one is shifted to right.
3. Otherwise, the element is inserted next to the compared item.

*/

package main

import "fmt"

func insertionSort(inputArray []int) []int {
	for i := 1; i <= len(inputArray)-1; i++ {
		temp := inputArray[i]
		j := i - 1
		for j >= 0 && inputArray[j] > temp {
			inputArray[j+1] = inputArray[j]
			j--
		}
		inputArray[j+1] = temp
	}
	return inputArray
}

func main() {
	var inputNumber, temporaryVariable int
	fmt.Println("Enter the value for number of elements:")
	fmt.Scanf("%d\n", &inputNumber)
	var userArray = make([]int, inputNumber)
	for i := 0; i < inputNumber; i++ {
		fmt.Printf("Enter the value of %d element:\n", i)
		fmt.Scanf("%d\n", &temporaryVariable)
		userArray[i] = temporaryVariable
	}
	fmt.Println("userArray:", userArray)
	fmt.Println("sorted array:", insertionSort(userArray))
}
