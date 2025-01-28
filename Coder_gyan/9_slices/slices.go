package main

import (
	"fmt"
	"slices"
)

// slice ->Dynamic array
func main(){
	// uninitialized slice is nill
	// var nums[]int

	// fmt.Println(nums)
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	// var nums= make([]int,2,5)
	// capacity- maximum no of elements can fit
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// nums[0]=3 
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// copy function

	// var nums= make([]int,0,5)
	// var nums2= make([]int,len(nums))
	

	// nums = append(nums,2)
	// fmt.Println(nums,nums2)

	// slice operator

	var nums=[]int{1,2,3}
	fmt.Println(nums[0:2])

	// slice package

	var nums1=[]int{1,2}
	var nums2=[]int{1,2}

	fmt.Println(slices.Equal(nums1,nums2))
}
