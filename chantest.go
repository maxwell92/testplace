package main
import (
    "log"
    "time"
    "sync"
)

var achan chan int
var bchan chan int

func sender(wg *sync.WaitGroup) {
    var i int
    var j int
    i = 1 
    j = 100

    for {
        achan <- i
        bchan <- j
        time.Sleep(time.Duration(1) * time.Second)
        i++
        j--
    }
}

func receiver(wg *sync.WaitGroup) {
// Deadlock
/*    for {
        select {
            case j := <-bchan :{
                i := <-achan
                log.Printf("%d -- %d\n", i, j)
            }
        }
    }
*/
    for {
        select {
            case i := <-achan :{
                j := <-bchan
                log.Printf("%d -- %d\n", i, j)
            }
        }
    }
}

func main() {
    achan = make(chan int)
    bchan = make(chan int)

    var wg sync.WaitGroup
    wg.Add(2)
    go sender(&wg)
    go receiver(&wg)
    wg.Wait()
}
