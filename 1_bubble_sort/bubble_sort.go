/*

Bubble Sort
1. Compare two adjacent elements.
2. If the first element is bigger than the second, swap them.
3. At the end of each cycle, the biggest item on the array moves to the end.
4. The next cycle through the array leave off the last positioned biggest item.
Time Complexity- O(n) in best case and O(n^2) in worst case
Space Complexity- O(n)

*/

package main

import "fmt"

func bubbleSort(inputArray []int) []int {
	for i := 0; i < len(inputArray)-1; i++ {
		flag := 0
		for j := 0; j < len(inputArray)-1-i; j++ {
			if inputArray[j] > inputArray[j+1] {
				temp := inputArray[j]
				inputArray[j] = inputArray[j+1]
				inputArray[j+1] = temp
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	return inputArray
}

func main() {
	var numberOfElements, temp int
	fmt.Println("Enter the value  of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var userInput = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element\n", i)
		fmt.Scanf("%d\n", &temp)
		userInput[i] = temp
	}
	fmt.Println("userInput", userInput)

	fmt.Println("Sorted Input Array:", bubbleSort(userInput))
}
