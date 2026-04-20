// Go program to illustrate send
// and receive operation
package main

import (
	"fmt"
	"sync"
)

// unbuffered channel
func myFunc(unbufChan <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-unbufChan)

	wg.Done()
}

func main() {

	//wg := sync.WaitGroup{}
	var wg sync.WaitGroup
	wg.Add(1)

	//channel initilization
	unbufferedChan := make(chan string)
	go myFunc(unbufferedChan, &wg)
	unbufferedChan <- "10"
	wg.Wait()
}
