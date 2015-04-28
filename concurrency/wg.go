package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
var wg sync.WaitGroup

func DoSomething(i int) {
	time.Sleep(time.Second)
	fmt.Println("Something", i)
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DoSomething(i)
	}
	wg.Wait()
}

//END OMIT
