package main

import(
	"fmt"
	"time"
	)
	
func main(){  

	ch := make(chan string)
	go worker("sheep", ch)  
	
	
	msg := <- ch  // receiving is a blocking operation
	// main() will wait untill worker() sends the message
	
	fmt.Println(msg)// sheep will be printed only once because main() stops
}

func worker(thing string,ch chan string){

	time.Sleep(time.Second*3)
	for i:=1;i<=5;i++{
		ch <- thing  // sending over channel is a blocking operation
		//worker() will wait untill main() is ready to recieve the message
		// Hence Attaining communication &  synchronization
		time.Sleep(time.Millisecond * 500)
	}
	
	
}
	

	