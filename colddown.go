package main
import (
    "log"
    "time"
    "sync"
)
var sigchan chan int
func sender(wg *sync.WaitGroup) {
    var i int
    i = 1
    for {
        sigchan<- i
        log.Printf("[sender]: %d\n", i)
        i++
        time.Sleep(time.Duration(1) * time.Second)
    }
    wg.Done()
}


func receiver(wg *sync.WaitGroup) {
    var i int
    i = 1
    for {
        select {
            case n := <-sigchan :{ 
                log.Printf("[receiver]: %d\n", n)
            }
        }
        time.Sleep(time.Duration(i) * time.Second)
        i++
    }
    wg.Done()
}


func main() {
    sigchan = make(chan int)
    var wg sync.WaitGroup
    wg.Add(2)
    go sender(&wg)
    go receiver(&wg)
    wg.Wait()
}
