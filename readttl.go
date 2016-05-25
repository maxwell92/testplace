package main
import (
    "fmt"
    "log"
    "sync"
    "os"
    "bufio"
    "time"
)

func readttl(wg *sync.WaitGroup) {
    fmt.Println("readttl is running...")
    //file, err := os.Open("/root/papertest/generator/ttl")
    //file, err := os.Open("/root/papertest/generator/rate")
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
        time.Sleep(time.Duration(5) * time.Second)
    }
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go readttl(&wg)
    wg.Wait()
}
