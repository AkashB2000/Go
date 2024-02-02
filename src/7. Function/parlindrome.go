package main

import(
	"fmt"
	)
	
func reverse(num int)int{
	
	fmt.Println("In Reverse Function")
	
	rev:=0
	rem:=0
	
	for num>0{
		
		rem=num%10
		rev=rev*10+rem
		num=num/10
	}
	return rev
}

func main(){

	n := 12345
	rev := reverse(n)
	
	fmt.Println(rev)
	
	if rev == n{
		fmt.Println(n," is a Parlindrome")
	} else{
		fmt.Println(n," is not a Parlindrome")
		}
		
	

}