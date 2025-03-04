/*

1. Iterate over the array to find the smallest element.
2. At the end of each cycle, this smallest element is swapped with the first one.
3. The next cycle excludes the last positioned smallest item.

*/

package main

import "fmt"

func main() {
	var arr = [6]int{7, 4, 10, 8, 3, 1}
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr)-1; j++ {
			//min := arr[j]
			if arr[j+1] < arr[j] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
