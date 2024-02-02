package main
 
import (
"fmt"
"time"
)
 
func main() {
	// Creating a buffered channel with a capacity of 3
	myChannel := make(chan int, 3)
	 
	// Sending values to the buffered channel
	go func() {
		for i := 1; i <= 5; i++ {
		myChannel <- i  // it will fill the buffer without waiting(1,2,3)
		//for sending 4,5 : it will wait for reciever to clear the buffer
		fmt.Printf("Sent: %d\n", i)
	}
	close(myChannel) // Closing the channel when done sending
	}()
	 
	// Receiving values from the buffered channel
	time.Sleep(time.Second*5) // Give the Goroutine some time to send values
	for value := range myChannel {
		fmt.Printf("Received: %d\n", value)
	}
}