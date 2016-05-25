package main
import (
    "time"
    "log"
    "fmt"
)

func main() {
    var sum int
    var count int
    ticker1 := time.NewTicker(time.Second * 5)
    go func() {
       for _ = range ticker1.C {
            log.Println(sum / count) 
            sum = 0
            count = 0
       } 
    }()
    for {
        sum += count
        count++
        time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(sum)
	fmt.Println(count)
    }
    ticker1.Stop()
}
