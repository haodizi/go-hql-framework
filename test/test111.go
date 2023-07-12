package main

import "fmt"

func main() {
	arr1 := make(map[int]int)
	arr1[1] = 1
	arr1[2] = 4
	arr1[3] = 9
	for _, i2 := range arr1 {
		fmt.Println(i2)
	}
	fmt.Println("exe...")
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
