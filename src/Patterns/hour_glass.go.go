package main

import(
	"fmt"
)

func main(){
	
	n:=7
	mid:=n/2  // mid = 3
	
	//upper half
	for i:=1;i<=mid;i++{

		for space:=1;space<i;space++{
			fmt.Print(" ")
		}
		for star:=0;star<=mid+1-i;star++{
			fmt.Print("* ")
		}
			
		fmt.Println()			
	}
	
	//lower half
	for j:=1;j<=mid+1;j++{
		
		for space:=1;space<=mid+1-j;space++{
			fmt.Print(" ")
		}
		for star:=0;star<j;star++{
			fmt.Print("* ")
		}
		
		fmt.Println()
	}
		
}