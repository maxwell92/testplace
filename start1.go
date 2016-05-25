package main
import (
    "log"
    "time"
    "sync"
    "os"
    "strconv"
    "os/exec"
)

var sigchan chan int

func scale(n int) {
//    cmd := exec.Command("./ScaleVmUp", strconv.Itoa(n))
//    cmd := exec.Command("./print", strconv.Itoa(n))
    
//    cmd := exec.Command("./print", "1")
//    cmd := exec.Command("/bin/sh", "-C", "print", "1")
    
    str, err := cmd.Output()
    if err != nil {
        log.Println(err)
    }
    
    log.Println(str)
    log.Println("[scale] running\n")
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
                } else {
                    go scale(n)
                    log.Printf("[scaler]: go scale(%d)\n", n)
                }
                log.Printf("[scaler]: %d\n", cnt)
            } 
        }
    }
}

func alerter() {
    for {
        time.Sleep(time.Duration(3) * time.Second)
        n := 1
        sigchan<- n  
        log.Printf("[alerter]: sent!\n")
    }
}

func main() {
    var wg sync.WaitGroup
    sigchan = make(chan int)
    wg.Add(2)
    go alerter()
    go scaler()
    wg.Wait() 
}
