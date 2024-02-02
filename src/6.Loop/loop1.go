package main

import(
	"fmt"
)

func main(){
	a := [5]int{1,2,3,4,5}
	
	for i :=0 ; i<5 ; i++{
		fmt.Println(i," ",a[i])
	}
	
	j:=0
	for j<5{
		fmt.Println(j)
		j++
	}
	
	
}
