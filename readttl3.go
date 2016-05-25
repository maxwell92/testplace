package main
import (
    "fmt"
    "log"
    "sync"
    "os"
    "bufio"
    "time"
    "strconv"
)

var timech chan float64

func readttl(wg *sync.WaitGroup) {
    fmt.Println("readttl is running...")
    file, err := os.Open("/root/papertest/autoscaler/ttlave")
    if err != nil {
        log.Println(err)
    }

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)
    
    for scanner.Scan() {
        str := scanner.Text()
        log.Println(str)
        n, _ := strconv.ParseFloat(str, 64)
        timech <- n
        log.Println(n)
        time.Sleep(time.Duration(5) * time.Second)
    }
    wg.Done()
}

func receiver(wg *sync.WaitGroup) {
    var t float64
    for {
        select {
            case t = <-timech :{
                log.Println("[receiver]: ")
                log.Println(t)
            }
        }
        if t > 0.33333 {
            log.Println("Overload!")
        }
    }
}

func main() {
    var wg sync.WaitGroup
    timech = make(chan float64)
    wg.Add(2)
    go readttl(&wg)
    go receiver(&wg)
    wg.Wait()
}
