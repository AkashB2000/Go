package main

import(
	"fmt"
)

func main(){
	vertices := make(map[string]int)
	
	vertices["triangle"]=3
	vertices["square"]=4
	vertices["hexagon"]=6
	
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	
	delete(vertices , "square")
	
	fmt.Println(vertices)
	
}

