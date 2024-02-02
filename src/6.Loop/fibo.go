package main

import(
	"fmt"
	)
	
func main(){

	n:=10
	
	first_term:=0
	second_term:=1
	
	fmt.Print(first_term," ",second_term," ")
	
	for i :=3;i<=n;i++{
		
		next := first_term + second_term
		first_term = second_term
		second_term = next
		
		fmt.Print(next," ")
	
	}
}
