package main
import (
    "log"
    "time"
    "sync"
    "os"
)

var sigchan chan int

func scale(n int) {
    cmd := exec.Command("./ScaleVmUp", strconv.Itoa(n))
    if _, err := cmd.Output(); err != nil {
        log.Println(err)
    }

}

func scaler() {
    var n int
    var cnt int = 0
    for {
        select {
            case n = <- sigchan: {
                cnt++
                if cnt == 2 {
                    log.Printf("[scaler]: cnt=2\n")
                    os.Exit(1)
                }
                log.Printf("[scaler]: %d\n", cnt)
                go scale(n)
            } 
        }
    }
}

func alerter() {
    time.Sleep(time.Duration(3) * time.Second)
    n := 1
    sigchan<- n  
    log.Printf("[alerter]: sent!\n")
}

func main() {
    var wg sync.WaitGroup
    sigchan = make(chan int)
    wg.Add(2)
    go alerter()
    go scaler()
    wg.Wait() 
}
