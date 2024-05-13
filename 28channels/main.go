package main

import (
	"fmt"
	"sync"
)

/*
Channels are a way so that multiple goroutines can talk to each other
channels are bidirectional, but make it clear we mark them with directions and make then unidirectional
*/

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2) 																	
	wg := &sync.WaitGroup{}

	// myCh <- 5 //feeding value in channel
	// fmt.Println(<-myCh) // receiving value from channel

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {  //read only
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)

		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { //write only

		// close(myCh)
		// myCh <- 6
		// myCh <- 5


		// close(myCh)


		myCh <- 6
		myCh <- 5
		close(myCh)

		wg.Done()
	}(myCh, wg)

	wg.Wait()

}

/*
fatal error: all goroutines are asleep - deadlock!

Deadlock, in the context of concurrent computing, refers to a situation where a set of processes are all blocked because they're each waiting for a resource
held by another process in the group. Imagine two people holding onto the ends of a jump rope – neither can jump until the other lets go of their end.
*/
