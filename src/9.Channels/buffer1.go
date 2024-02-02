package main

import(
	"fmt"
	)


func main(){
	
	ch := make(chan string,2) // we can fill a Buffered Channel without a corresponding reciever and it wont block untill channel is full
	
	ch <- "hello" // we can send 2 messages over the channel without blocking
	ch <- "world"
	
	//ch <- "good morning"   can't send more messages because there are no reciever waiting
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
		
	
}
}

