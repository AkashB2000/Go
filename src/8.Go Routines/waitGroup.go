package main

import(
	"fmt"
	"time"
	"sync"
	)
	
func main(){  

	var wg sync.WaitGroup
	wg.Add(1)  // main() now have 1 go routine to wait for
	go worker("sheep", &wg)  
	
	for i:=1;i<=10;i++{
		fmt.Println("main ",i)
	}
		
	wg.Wait() // main() will wait untill wg=0
	
}

func worker(thing string, wg *sync.WaitGroup){
	defer wg.Done()  // decrements counter by 1
	//when a go routine completes its task , it calls wg.Done() to signal that it has finished its work
	//defer : schedules wg.done() to run after worker() stops
	// no matter how worker() exits , wg.Done() will be called
	for i:=1;i<=5;i++{
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	
	// wg.Done() : if worker() exits prematurely , wg.Done() is never called and main() keeps waiting, resulting in a Deadlock
}
	