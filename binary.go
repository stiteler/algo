package main 

import (
	"fmt"
	"github.com/steaz/algo"
)

func main() {
	hay := []int{3, 4, 2, 34, 2, 5, 4, 5, 6, 7, 8, 9, 6, 5, 6, 7, 7, 6 , 5, 4, 3, 54}
	//sort it (test mergesort)
	algo.Mergesort(hay)
	fmt.Println(hay)

	ndl := 1 //won't be found
	ndl2 := 2 //will be found

	
	fmt.Println(BinarySearch(ndl, hay))
	fmt.Println(BinarySearch(ndl2, hay))

}
// slice binary search, saves recursive memory, each call takes up logarithmically less space
// than the normal version which passes the whole array. Wonder how much time it takes to 
// slice the slice? 
func BinarySearch(ndl int, hay []int) bool {
	//calculate mid everytime. 
	mid := int(len(hay)/2)
	//we need thhis because 
	if(mid < 1) { return false }

	if ndl == hay[mid] {
		return true
	} else if ndl < hay[mid] {
		return BinarySearch(ndl, hay[:mid])
	} else if ndl > hay[mid] {
		return BinarySearch(ndl, hay[mid:])
	} else {
		return false
	}
}