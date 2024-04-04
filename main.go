package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0, 3, 7, 35, 5, 10, 9}
	sortColors(nums)
	fmt.Println(nums)
}
func sortColors(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quicksort(arr, left, pivot-1)
		quicksort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
