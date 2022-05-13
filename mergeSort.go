//merge sort is divide and conquer
//divides input array in two halves
//calls itself for two halves
//merges two sorted halves
//merge(arr,left,middle,right) assumes that arr[left:middle] and arr[middle+1:right] are sorted
package main

import "fmt"

func merge(anArray []int, left int, middle int, right int) {
	subArrayOne := middle - left + 1
	subArrayTwo := right - middle
	leftArray := make([]int, subArrayOne)
	rightArray := make([]int, subArrayTwo)

	for i := 0; i < subArrayOne; i++ {
		leftArray[i] = anArray[left+i]
	}
	for i := 0; i < subArrayTwo; i++ {
		rightArray[i] = anArray[middle+1+i]
	}

	indexOfSubarrayOne := 0
	indexOfSubarrayTwo := 0
	indexOfMergedArray := left

	for indexOfSubarrayOne < subArrayOne && indexOfSubarrayTwo < subArrayTwo {
		if leftArray[indexOfSubarrayOne] <= rightArray[indexOfSubarrayTwo] {
			anArray[indexOfMergedArray] = leftArray[indexOfSubarrayOne]
			indexOfSubarrayOne++
		} else {
			anArray[indexOfMergedArray] = rightArray[indexOfSubarrayTwo]
			indexOfSubarrayTwo++
		}
		indexOfMergedArray++
	}
	//copy remaining elements of left if there are any
	for indexOfSubarrayOne < subArrayOne {
		anArray[indexOfMergedArray] = leftArray[indexOfSubarrayOne]
		indexOfSubarrayOne++
		indexOfMergedArray++
	}

	for indexOfSubarrayTwo < subArrayTwo {
		anArray[indexOfMergedArray] = rightArray[indexOfSubarrayTwo]
		indexOfSubarrayTwo++
		indexOfMergedArray++
	}
}

func mergesort(anArray []int, begin int, end int) {
	if begin < end {
		mid := begin + (end-begin)/2
		mergesort(anArray, begin, mid)
		mergesort(anArray, mid+1, end)
		merge(anArray, begin, mid, end)
	}
}

func main() {
	anArray := []int{12, 11, 13, 5, 6, 7}
	mergesort(anArray, 0, len(anArray)-1)
	for i := 0; i < len(anArray); i++ {
		fmt.Println(anArray[i])
	}
}
