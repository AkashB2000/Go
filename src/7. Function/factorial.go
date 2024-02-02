package main

import(
	"fmt"
	)
	
func fact(n int)int{
	
	f:=1
	for i:=1;i<=n;i++{
		f=f*i
		
	}
	
	return f
}

func recursion(n int)int{
	if n==0 || n==1{
		return 1
	}
	
	return n*recursion(n-1)
	
}

func main(){
	n:=6
	
	result1 := fact(n)
	
	result2 := recursion(n)
	
	fmt.Println(result1," ",result2)
}