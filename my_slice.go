// slice_test.go
package main

import (
	"fmt"
)

func printMySlice(sli []int) {
	fmt.Println()
	for _, v := range sli {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Element of array: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElement of slice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	mySlice1 := make([]int, 5, 20)
	for i, _ := range mySlice1 {
		mySlice1[i] = i
	}
	for _, v := range mySlice1 {
		fmt.Println(v, " ")
	}

	fmt.Println("mySlice1's lenth:", len(mySlice1), ", cap:", cap(mySlice1))

	mySlice1 = append(mySlice1, 1, 2, 3)
	printMySlice(mySlice1)
	mySlice1 = append(mySlice1, mySlice...)
	printMySlice(mySlice1)

	mySlice2 := []int{11, 12, 13, 14, 15}
	mySlice3 := []int{3, 4, 5, 6}
	printMySlice(mySlice3)
	copy(mySlice3, mySlice2)
	printMySlice(mySlice3)
	copy(mySlice1, mySlice2)
	printMySlice(mySlice1)

	fmt.Println()
}
