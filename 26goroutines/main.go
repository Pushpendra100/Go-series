package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

// https://www.youtube.com/watch?v=V-0ifUKCkBI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=52
/*
=> Concurrency vs Parallelism

Concurrency - you are allowing yourself to involve (hop from one task to another) in other tasks while you are already doing something but you are not doing them at the same time
Parallelism - you are doing multiple things at the same time


=> Goroutines - achieve parallelism

		Thread						Goroutines
	Managed by OS				Managed by Go runtime
	Fixed stack - 1MB			Flexible stack - 2KB

Do not communicate by sharing memory; instead share memory by communicating

*/

// func main(){
// 	go greeter("Hello")
// 	greeter("World")
// }

// func greeter(s string){
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

/*
Hello
World
World
Hello
World
Hello
World
Hello
World
Hello
World
Hello
*/



// ******************* Wait Groups ******************************************
/*
wait groups - They are useful when you need to wait for multiple goroutines to finish their tasks before proceeding with the main program.
*/

var signals = []string{"test"}
var wg sync.WaitGroup // usually these are pointers
var mut sync.Mutex //pointer

func main(){
	websitelist := []string{
		"https://youtube.com",
		"https://linkedin.com",
		"https://google.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode (endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}


// MUTEX - A mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex	
/*
In computer science, a mutex, or mutual exclusion object, is a synchronization primitive that prevents multiple threads from accessing the same shared resource at once. A shared resource is a code element with a critical section, which is the part of the code that should not be executed by more than one thread at a time. 
Mutex is a synchronization primitive that grants exclusive access to the shared resource to only one thread. If a thread acquires a mutex, the second thread that wants to acquire that mutex is suspended until the first thread releases the mutex.
*/