package main

import "fmt"

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 3}
	fmt.Printf("%f\n", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	pivot := (n / 2)

	var i int
	var j int

	var res []int
	for len(res) <= pivot {
		if i < len(nums1) && (j == len(nums2) || nums1[i] <= nums2[j]) {
			res = append(res, nums1[i])
			i++
		}

		if j < len(nums2) && (i == len(nums1) || nums1[i] > nums2[j]) {
			res = append(res, nums2[j])
			j++
		}
	}

	if n&1 == 0 {
		return float64(res[pivot])/2 + float64(res[pivot-1])/2
	} else {
		return float64(res[pivot])
	}
}
