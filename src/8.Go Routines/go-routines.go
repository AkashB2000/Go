package main

import(
	"fmt"
	"time"
	)
	
func main(){  // main() is also a routine
	go count("sheep")  // a go routine is created and run parallely to main routine
	
	for i:=1;i<=10;i++{
		fmt.Println("main ",i)
		
	}
	// all routines stops if main() comes to end
	fmt.Scanln()  // to keep main() alive so that count() routine can run
	
}

func count(thing string){
	for i:=1;true;i++{
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	
}
	

	