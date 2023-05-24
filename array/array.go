package main

import "fmt"

func avg(arr []int) bool {
	x := 0
	sum := 0
	//a:=0
	for i := 0; i < len(arr); i++ {
		x = arr[i]
		//n:=len(arr)
		for x > 9 {
			x = x / 10
		}
		sum = x + (sum * 10)
	}
	//fmt.Println(sum)
	if (sum % 7) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Printf("Please enter the size of an array: ")
	var n int
	fmt.Scanln(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("enter %dth an element: ", i)
		fmt.Scanln(&arr[i])
	}
	res := avg(arr)
	fmt.Println(res)
}
