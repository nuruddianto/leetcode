package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{3, 2, 1}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	stay := nums1
	move := nums2

	if len(nums2) > len(nums1) {
		stay = nums2
		move = nums1
	}

	var res []int
	var i int
	var j int
	for i < len(move) {
		for j < len(stay) {
			if move[i] == stay[j] {
				res = append(res, move[i])
				stay = append(stay[:j], stay[j+1:]...)
				//fmt.Println(stay)
				j = 1000
			}
			j++
		}
		j = 0
		i++
	}

	return res
}
