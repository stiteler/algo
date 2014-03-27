package main

import (
	"fmt"
)

func main() {
	unsorted := []int{5, 3, 7, 0, 13, 53, 44, 66, 75, 86, 23, 4, 3, 53}
	Mergesort(unsorted)
	fmt.Println(unsorted)

}

func Mergesort(in []int) {
	temp := make([]int, len(in))
	sort(in, temp, 0, len(in)-1)
}

func sort(in, temp []int, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	sort(in, temp, lo, mid)
	sort(in, temp, mid+1, hi)
	merge(in, temp, lo, mid, hi)
}

func merge(in, temp []int, lo, mid, hi int) {

	for i := lo; i <= hi; i++ {
		temp[i] = in[i]
	}

	left := lo
	right := mid + 1

	for k := lo; k <= hi; k++ {
		if left > mid {
			in[k] = temp[right]
			right++
		} else if right > hi {
			in[k] = temp[left]
			left++
		} else if temp[left] < temp[right] {
			in[k] = temp[left]
			left++
		} else {
			in[k] = temp[right]
			right++
		}
	}
}
