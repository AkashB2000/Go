package main

import(
	"fmt"
)

func main(){

	student := map[string]int{
		"akash" : 90,
		"prince" : 100,
		"tishant" : 100,
		"ashutosh" : 80,
	}
	
	for key,value := range student{   // range is used to iterate over elements of array , slice, string , map , channel
		fmt.Printf("Name : %s  Marks : %d\n" , key , value)	
	}
	
	arr := [5]int{10,20,30,40,50}
	
	for index,value := range arr{
		fmt.Printf("Index : %d  Value : %d\n", index,value)
	}
	
	str:="hello world"
	
	for index, char := range str{
		fmt.Printf("Index : %d  Character : %c\n" , index , char)
	}
}


