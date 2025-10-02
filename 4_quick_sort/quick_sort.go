/*

1. It uses divide and conquer.
2. A pivot element is choosen, determining its final position in the array.
3. After that smaller elements are on the left of the pivot element and larger elements are on the right of the pivot element.
4. The left and right sub-arrays are processed recursively in the same way.
Time Complexity- O(n log (n)) in best and in average case and O(n ^ 2) in worst case.
Space Complexity- O(n)

*/

package main

import "fmt"

func quickSort(inputArray []int, lbound, ubound int) []int {
	if lbound < ubound {
		inputArray, location := partition(inputArray, lbound, ubound)
		quickSort(inputArray, lbound, location-1)
		quickSort(inputArray, location+1, ubound)
	}
	return inputArray
}

func partition(inputArr []int, lowerBound, upperBound int) ([]int, int) {
	pivotElement := inputArr[lowerBound]
	start := lowerBound
	end := upperBound - 1
	for start < end {
		for inputArr[start] <= pivotElement {
			start++
		}
		for inputArr[end] > pivotElement {
			end--
		}
		if start < end {
			inputArr[start], inputArr[end] = inputArr[end], inputArr[start]
		}
	}
	if start > end {
		inputArr[lowerBound], inputArr[end] = inputArr[end], inputArr[lowerBound]
	}
	return inputArr, end
}

func main() {
	var numberOfElements, temporaryVariable int
	fmt.Println("Enter the value of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var userInputArray = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element:\n", i)
		fmt.Scanf("%d\n", &temporaryVariable)
		userInputArray[i] = temporaryVariable
	}
	fmt.Println("user input array:", userInputArray)
	lb, ub := 0, numberOfElements-1
	fmt.Println("Sorted array using quick sort mechanism:", quickSort(userInputArray, lb, ub))
}
