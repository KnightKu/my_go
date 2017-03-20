// my_closure
package main

import (
	"fmt"
	"log"
)

func main() {
	var j int = 5

	f := func() func() {
		i := 10
		return func() {
			j += 1 //这里是引用，不是行参，改了j就是对外围的j生效
			fmt.Printf("i[%d], j[%d]\n", i, j)
		}
	}()

	f()

	j *= 2
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error:%v", r)
		}
	}()

	f()

}
