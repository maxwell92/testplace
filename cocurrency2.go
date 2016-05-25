// Maintainer: liyao.miao@yeepay.com

package main
import (
	"fmt"
	"time"
	"sync"
)

func talk(wg *sync.WaitGroup) {
	for {
		fmt.Println("Routine1 talking...")
		time.Sleep(time.Duration(1) * time.Second)
	}
	wg.Done()
}

func speak(wg *sync.WaitGroup) {
	for {
		fmt.Println("Routine2 speaking...")
		time.Sleep(time.Duration(2) * time.Second)
	}
	wg.Done()
}

func tell(wg *sync.WaitGroup) {
	for {
		fmt.Println("Routine3 telling...")
		time.Sleep(time.Duration(3) * time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for {
			talk(&wg)					
		}
	}()
	go func() {
		for {
			speak(&wg)	
		}
	}()
        go func() {
		for {
			tell(&wg)
		}	
	}()
	wg.Wait()
	//time.Sleep(time.Duration(10) * time.Second)
}
