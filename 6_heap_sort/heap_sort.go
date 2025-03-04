package main

import "fmt"

func linearSort(inputArray []int) {
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == 49 {
			fmt.Printf("Match Found, Element exists at index: %d\n", i)
			// return i
		} else {
			fmt.Println("Match not found, element doesnot exists")
		}
	}
	//fmt.Println("Match not found, Element doesnot exists")

}

func main() {
	var numberOfElements, temp int
	fmt.Println("Enter the value of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var arr = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element:\n", i)
		fmt.Scanf("%d\n", &temp)
		arr[i] = temp
	}
	fmt.Println("array:", arr)
	// fmt.Println("finding the presence of element at index:", linearSort(arr))
	linearSort(arr)

}
