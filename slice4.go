package main
import "fmt"

func main() {
    var win []int
    win = make([]int, 0, 10)
    
    win = append(win, 1)
    win = win[1: 10]
    fmt.Println(win)

    win = append(win, 2)
    win = win[1: 10]
    fmt.Println(win)

    win = append(win, 3)
    win = win[1: 10]
    fmt.Println(win)

    win = append(win, 4)
    win = win[1: 10]
    fmt.Println(win)

}
