package main
import (
    "log"
    "time"
)

func main() {
    sigChan := make(chan int)
    go func() {
	for {
            select {
                case signal := <- sigChan: {
  		    if signal == 1 {

		        log.Println("Mushroom")
                    } 
                } 
            }
        }

    }()

    for i := 0; i < 20; i++ {
        time.Sleep(time.Duration(1) * time.Second)
	sigChan <- i % 3
    }
}
