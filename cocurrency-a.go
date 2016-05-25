// WARNING: wrong code
// Maintainer: maxwell92@126.com
package main
import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func(&wg) {
		for {
			fmt.Println("Routine1 talking...")
			time.Sleep(time.Duration(1) * time.Second)
		}
		wg.Done()
	}(wg *sync.WaitGroup)
	go func(&wg) {
		for {
			fmt.Println("Routine2 speaking...")
			time.Sleep(time.Duration(2) * time.Second)
		}
		wg.Done()
	}(wg *sync.WaitGroup)
        go func(&wg) {
		for {
			fmt.Println("Routine3 telling...")
			time.Sleep(time.Duration(3) * time.Second)
		}	
		wg.Done()
	}(wg *sync.WaitGroup)
	wg.Wait()
	//time.Sleep(time.Duration(10) * time.Second)
}
