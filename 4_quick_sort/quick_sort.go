/*

1. It uses divide and conquer.
2. A pivot item is choosen, determining its final position in the array.
3. After that smaller items are on the left of the pivot and larger are on the right.
4. The left and right sub-arrays are processed recursively in the same way.

*/

package main

import "fmt"

func quickSort(inputArray []int) {}

func main() {
	var numberOfElements, temporaryVariable int
	fmt.Println("Enter the value of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var arr = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element:\n", i)
		fmt.Scanf("%d\n", &temporaryVariable)
		arr[i] = temporaryVariable
	}
	fmt.Println("arr:", arr)
}
