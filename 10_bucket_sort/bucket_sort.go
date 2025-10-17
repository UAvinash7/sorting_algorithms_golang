/*

Bucket Sort or Bin Sort works by distributing the elements of an array into a number of buckets.

Each bucket is then sorted individually, either using a different sorting algorithm or by recursively applying the bucket sorting algorithm.

It works as-

1. Setup an array of initially empty "buckets".
2. Scatter: Go over the original array, putting each element in its bucket.
3. Sort each non-empty bucket.
4. Gather: Visit the buckets, in order and put all elements back into the original array.

*/

package main

import (
	"fmt"
	"sort"
)

// bucketSort sorts a slice of floats using the bucket sort algorithm.
func bucketSort(arr []float64) []float64 {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// Find the maximum value to determine the range for buckets
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	// Determine the number of buckets. A common choice is sqrt(n) or n/k.
	// For simplicity, we'll use n buckets for now.
	numBuckets := n
	if numBuckets == 0 { // Avoid division by zero if arr is empty after initial check
		return arr
	}

	// Create buckets as a slice of slices of floats
	buckets := make([][]float64, numBuckets)

	// Distribute elements into buckets
	for _, val := range arr {
		// Calculate bucket index. Ensure index is within bounds.
		// Scale the value to fit into the bucket indices [0, numBuckets-1]
		bucketIndex := int((val / (maxVal + 1)) * float64(numBuckets)) // maxVal + 1 to handle maxVal itself
		buckets[bucketIndex] = append(buckets[bucketIndex], val)
	}

	// Sort each bucket and concatenate them
	result := make([]float64, 0, n)
	for _, bucket := range buckets {
		sort.Float64s(bucket) // Sort each bucket using Go's built-in sort
		result = append(result, bucket...)
	}

	return result
}