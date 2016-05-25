package main
import (
    "fmt"
    "time"
)

func main() {
    var t float32
    for {
        if t < 5.0 {
            t++
        } 
        fmt.Println(t)
        time.Sleep(time.Duration(1) * time.Second)
    }
}
