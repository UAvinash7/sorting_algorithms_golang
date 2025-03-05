/*

1. It uses divide and conquer approach.
2. The array is divided into two equal halves. Each half is recursively processed.
3. At the end off each recursive step, there are two sorted sub-arrays.
4. These two sorted sub-arrays are merged into a sorted one.

*/

package main

func linearSearch(){

}

func main() {
	var numberOfElements, temporaryVariable int
	fmt.Println("Enter the value of number of elements:")
	fmt.Scanf("%d\n", &numberOfElements)
	var arr = make([]int, numberOfElements)
	for i := 0; i < numberOfElements; i++ {
		fmt.Printf("Enter the value of %d element\n")
		fmt.Scanf("%d\n", &temporaryVariable)
	}
}

