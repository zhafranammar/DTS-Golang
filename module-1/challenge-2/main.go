package main

import (
	"fmt"
)

func main() {
	s := "САШАРВО"
	for i:=0; i<=10;i++{
		if i==5{
			for i, v := range s {
				fmt.Printf("character %U %q starts at byte position %d\n",v,v,i)
			}		
		}else if  i<5 {
			fmt.Printf("Nilai i = %d\n", i)
		}else{
			fmt.Printf("Nilai j = %d\n", i)
		}
	}
}