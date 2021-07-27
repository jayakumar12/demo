package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//m:=make(map[int64]string)
	//fmt.Print("generated  random number: ")
	//Random := rand.Intn(1000000)
	//fmt.Println(Random)
	m := make(map[int]string)
	for i := 0; i < 100; i++ {

		fmt.Printf(m[rand.Intn(100000000)])
		m[rand.Intn(100000000)] = "jhauekiauert"
	}
	//m[rand.Intn(1000000)] = "jhauekiauert"
	fmt.Println(m)
	fmt.Println("enter the number u want to search::")
	var input int
	fmt.Scanln(&input)
	v, ok := m[input]
	fmt.Println("The value:", v, "Present?", ok)

}
