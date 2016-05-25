package main
import (
    "fmt"
    "log"
    "sync"
)

func monitorttl(wg *sync.WaitGroup) {
     
}

func monitorrate(wg *sync.WaitGroup) {

}

func monitorave(wg *sync.WaitGroup) {

}

func monitor() {
    var wg sync.WaitGroup	
    wg.Add(3)
    go monitorttl(&wg)
    go monitorrate(&wg)
    go monitorave(&wg)
    wg.Wait()
}

func main() {
    monitor()  
}
