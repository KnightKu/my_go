// my_map.go
package main

import (
	"fmt"
)

func printPerson1(args ...Personinfo) {
	for _, person := range args {
		fmt.Println(person.Id, person.Name)
	}
}

func printPerson(args ...Personinfo) {
	printPerson1(args[:1]...)
}

type Personinfo struct {
	Id   string
	Name string
}

func main() {
	var personDB map[string]Personinfo
	personDB = make(map[string]Personinfo, 20)

	personDB["123"] = Personinfo{"123", "Lucy"}
	personDB["124"] = Personinfo{"124", "Tom"}

	person, ok := personDB["125"]
	if ok {
		fmt.Println("Get person: ", person.Name, person.Id, ", for key:125")
	} else {
		fmt.Println("Can not found person for id:125")
	}
	person, ok = personDB["123"]
	if ok {
		fmt.Println("Get person: ", person.Name, person.Id, ", for key:123")
	} else {
		fmt.Println("Can not found person for id:123")
	}
	delete(personDB, "124")
	personDB["126"] = Personinfo{"126", "Cengku"}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	for {
		sum++
		if sum > 100 {
			break
		}
	}
level0:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break level0
			}
			fmt.Println("i:", i, " j:", j)
		}
	}
	fmt.Println(sum)
	printPerson(personDB["126"], personDB["123"])

}
