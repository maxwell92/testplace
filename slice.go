package main
import (
    "fmt"
)

func main() {
    hWin := make([]float64, 0, 10)
    //hWin[0] = 1.0
    //fmt.Println(hWin)
    fmt.Println(len(hWin))
    hWin = append(hWin, 1.0)
    fmt.Println(hWin)
    fmt.Println(len(hWin))
    fmt.Printf("xxx: %f\n", hWin[0])
    
}
