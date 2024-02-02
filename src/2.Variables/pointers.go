package main

import(
	"fmt"
	)
	
func main(){

	x:=5
	fmt.Println(&x)
	
	x=10
	fmt.Println(&x)  // address doesn't change
	
	p:=&x
	fmt.Println(*p)  // de-referencing a pointer
	
	
	
	
	
}
