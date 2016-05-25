package main
import (
    "log"
    "sync"
    "time"
)

var achan chan int
var bchan chan int

func sender(wg *sync.WaitGroup) {
    i := 1
    j := 100
    for { 
        achan <- i
        bchan <- j 
        i++
        j++
        time.Sleep(time.Duration(2) * time.Second)
    }
    wg.Done()
}

func receiver(wg *sync.WaitGroup) {
    var i int
    var j int
    for {
        select {
            case i = <-achan :{
                log.Printf("i: %d\n", i) 
            }
            case j = <-bchan :{
                log.Printf("j: %d\n", j)
            }
       }

        if i > 10 {
            log.Printf("i: %d > 10\n", i)
        }
    
        if j < 85 {
            log.Printf("j: %d < 85\n", j)
        }

        if i > 5 && j < 93 { 
            log.Printf("i: %d > 5 && j: %d < 93\n", i, j)
        }
    }
    wg.Done()
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
