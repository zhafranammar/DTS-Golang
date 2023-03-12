package main

import (
	"fmt"
)

func main() {
	i:=21
	j:=true
	k:= 123.456
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	fmt.Println(j)
	fmt.Printf("%U\n", []rune("Я"))
	fmt.Printf("%d\n", 21)
	fmt.Printf("%o\n", 21)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%X\n", 15)
	fmt.Printf("%f\n", k)
	fmt.Printf("%E\n", k)
}