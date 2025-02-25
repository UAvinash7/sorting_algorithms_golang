/*

1. Iterate over the array to find the smallest element.
2. At the end of each cycle, this smallest element is swapped with the first one.
3. The next cycle excludes the last positioned smallest item.

*/

package main

import "fmt"

func main() {
	var arr = [6]int{7, 4, 10, 8, 3, 1}
	var arrResult []int
	var smallestNumber int
	for i := 0; i < len(arr)-1; i++ {
		//smallestNumber := arr[i]
		for j := i; j < len(arr)-1; j++ {
			smallestNumber = arr[j]
			if arr[j] > arr[j+1] {
				smallestNumber = arr[j+1]
			}
		}
	}
	arrResult = append(arrResult, smallestNumber)
	fmt.Println("RESULT:", arrResult)
}
