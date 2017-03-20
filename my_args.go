// my_args.go
package main

import (
	"fmt"
)

func myPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is unknow type")
		}
	}
	// 匿名函数
	f := func(len int) int {
		return len
	}
	fmt.Println(f(len(args)))
}

func main() {
	v1 := 123
	v2 := "xyz"
	v3 := int64(345)
	v4 := 1.23
	myPrint(v1, v2, v3, v4)
	fmt.Println("Hello World!")
}
