package main
import (
    "log"
    "time"
    "sync"
)

var isFreezen int
var sigChan chan int

func sender(wg *sync.WaitGroup) {
    n := 1
    for {
        log.Printf("[sender]: %d\n", n)
        sigChan <- n
        n++
        time.Sleep(time.Duration(1) * time.Second)
    } 
    wg.Done()
}

func freezer(wg *sync.WaitGroup, t int) {
    ticker1 := time.NewTicker(time.Second * time.Duration(t))
    for {
        if isFreezen == 1 {
            <-ticker1.C
            isFreezen = 0 
        } else {
            time.Sleep(time.Duration(1) * time.Second)
        } 
    }
    wg.Done() 
}

func receiver(wg *sync.WaitGroup) {
    for {
        select {
            case n := <- sigChan : {
                if isFreezen == 0 {
                    log.Printf("[receiver]: %d\n", n)     
                    isFreezen = 1         
                } else {
                    log.Printf("[receiver]: %d but it's freezen!\n", n)
                }
            }
        }
    } 
    wg.Done()
}

func main() {
    isFreezen = 0
    
    sigChan = make(chan int)

    var wg sync.WaitGroup
    wg.Add(3)
    go sender(&wg)
    go receiver(&wg)
    go freezer(&wg, 3)
    wg.Wait()
}
