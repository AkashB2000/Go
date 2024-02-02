package main

import(
	"fmt"
	"time"
	)
	
func main(){  

	ch := make(chan string)
	go worker("sheep", ch)  
/*
	for{ // to recieve all message from sender, recive message untill channel is open
		msg, open:= <- ch 
		if !open{  // if channel is not open , dont wait for recieving msg
			break
		}
		fmt.Println(msg)
	}

*/

	for msg:= range ch {  // better way . No need to manually check if channel is closed

		fmt.Println(msg)
		time.Sleep(time.Second*5)
	}

}

func worker(thing string,ch chan string){
	 
	for i:=1;i<=5;i++{
		ch <- thing  
		time.Sleep(time.Millisecond * 500)
	}
	
	close(ch) // only sender should close the channel
	
}

	