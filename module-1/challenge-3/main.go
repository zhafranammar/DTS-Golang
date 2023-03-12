package main

import (
	"fmt"
)


func solve(s string) {
	var charMap = map[string]int{}
	for _, v := range s {
		fmt.Printf("%s\n", string(v))
		charMap[string(v)]++
	}
	fmt.Printf("%v\n", charMap)
}

func main() {
	solve("Selamat malam") // Ubah Kalimatnya disini
}