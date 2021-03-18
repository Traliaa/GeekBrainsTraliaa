package main

import "fmt"

func main() {

	arr := []int{3, 232, -1, 34, 6565, 43434}

	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int) {
	for k, v := range arr {
		if k != 0 {
			for k >= 1 && arr[k-1] > v {
				arr[k] = arr[k-1]
				k--
			}
			arr[k] = v
		}
	}
}
