package main

import "fmt"

func find_first_duplicate(arr []int) int {
	checked := []int{}
	for i := 0; i < len(arr); i++ {
		for _, val := range checked {
			if val == arr[i] {
				return arr[i]
			}
		}

		checked = append(checked, arr[i])
	}
	return -1
}

func main() {
	fmt.Println(find_first_duplicate([]int{1, 2, 3, 3, 0}))
}
