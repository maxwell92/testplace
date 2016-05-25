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

var ttlch chan float64
var ratech chan float64

func readrate(wg *sync.WaitGroup) {
    fmt.Println("readrate is running...")
    file, err := os.Open("/root/papertest/generator/rate")
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
        ratech <- n
        log.Println(n)
        time.Sleep(time.Duration(5) * time.Second)
    }
    wg.Done()
    log.Println("readrate exit....")
}



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
        ttlch <- n
        log.Println(n)
        time.Sleep(time.Duration(5) * time.Second)
    }
    wg.Done()
    log.Println("readttl exit....")
}

func receiver(wg *sync.WaitGroup) {
    for {
        select {
            case t := <-ttlch :{
                log.Println("[ttl]: ")
                log.Println(t)
            }
            case r := <-ratech :{
                log.Println("[rate]: ")
                log.Println(r)
            }
        }
    }
}

func main() {
    var wg sync.WaitGroup
    ttlch = make(chan float64)
    ratech = make(chan float64)
    wg.Add(3)
    go readttl(&wg)
    go readrate(&wg)
    go receiver(&wg)
    wg.Wait()
}
