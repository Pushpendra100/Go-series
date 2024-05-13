package main

import (
	"fmt"
	"sync"
)

// using this command check for any race conditions are there or not - we want to remove race conditions -- use mutex
//  go run --race .

func main(){
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}

	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}
	// for write we can use simply Mutex, but for reading ability also we use RWMutex
	// read is fast in comparison to write, thus lock is imp for writing

	var score = []int{0}

	wg.Add(4)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("One goroutine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Two goroutine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Three goroutines")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Fourth testing goroutines")
		m.RLock()
		fmt.Println("score", score)
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()  // --> to give all the goroutines to comeback and report

	fmt.Println(score)
}


// To stop data race we use mutex