package main

import "fmt"

func linearSort(inputArray []int) int {
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
	fmt.Println("finding the presence of element", linearSort(arr))

}
