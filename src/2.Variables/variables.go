package main

import(
	"fmt"
)

func main(){
	var x int // default value 0 for int types
	fmt.Println(x)  // must use variable after declaring
	
	var y int=7
	z:=5  // go can infer that z is an int type
	sum:=y+z
	
	fmt.Println(sum)
	
	addressOfx := &x  
	fmt.Printf("Value : %d, Address : %p\n", x, addressOfx)  // %p = memory in hexadecimal format
	
}