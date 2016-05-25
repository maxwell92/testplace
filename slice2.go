package main
import (
    "fmt"
    "time"
)

var hWin []float64

func addToSlice(win []float64, size int, n float64) {
	win = append(win ,n)
	win = win[1: size]
}

func slice() {
    for i := 1.0; i < 20.0; i++ {
        fmt.Println(hWin)
        fmt.Println(len(hWin))
        //addToSlice(hWin, 10, i)
        hWin = append(hWin ,i)
	    hWin = hWin[1: 10]
        time.Sleep(time.Duration(1) * time.Second)
    }  
}

func main() {
    hWin = make([]float64, 0, 10)
    go slice()     
    time.Sleep(time.Duration(1000) * time.Second) 
}
