package main

import(
	"fmt"
)

func main(){

	student := map[string]int{  // no need of make if map is initialised
		"akash" : 90,
		"prince" : 100,
		"tishant" : 100,
		"ashutosh" : 80,
	}
	
	for key,value := range student{
		fmt.Printf("Name : %s  Marks : %d\n" , key , value)
	}
	
}
