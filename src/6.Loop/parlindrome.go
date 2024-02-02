package main

import(
	"fmt"
)

func main(){

	n := 1232134
	temp :=n
	
	rev:=0
	
	rem:=0
	
	for n>0{
		rem=n%10
		rev=rev*10+rem
		n=n/10
	}
	
	if rev==temp{
		fmt.Println(temp," is a parlindrome no")
	}else{
			fmt.Println(temp," is not a parlindrome no")
		}
		
	}