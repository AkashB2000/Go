package main

import(
	"fmt"
	"time"
	)
	
func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func(){
		for{
			c1<-"Every 500ms"
			time.Sleep(time.Millisecond*500)
		}
	}()
	
	go func(){
		for{
			c2<-"Every 2 seconds"
			time.Sleep(time.Second*2)
		}
		}()
/*		
	for{
		fmt.Println(<-c1) Not ideal because c1 has to wait for c2 for 2s every time
		2nd Go Routine is slowing down the first Go Routine
		fmt.Println(<-c2) We should get messages from c1 independent of c2
		}
*/
	for{
		select{
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			}
		}

		
	}

	