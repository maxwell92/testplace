package main
import (
    "log"
    "time"
    "os"
)

func worker() {
    var i int
    i = 0
    for {
        log.Printf("[worker]: %d\n", i)
        i++
        time.Sleep(time.Duration(1) * time.Second)
        if i == 10 {
            os.Exit(-1)
        }
    }
}

func main() {
    go worker()
    var i int
    i = 0
    for {
        log.Printf("[main]: %d\n", i)
        i++
        time.Sleep(time.Duration(1) * time.Second)
        if i == 30 {
            os.Exit(-2)
        }
    }
      
}
