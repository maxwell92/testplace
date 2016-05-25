package main
import "fmt"

func main() {
    var total []int
    total = make([]int, 10)
    fmt.Println(total)
    fmt.Println(len(total))
    total[0] = 1
    fmt.Println(total)
    fmt.Println(len(total))
    total[1] = 2
    fmt.Println(total)
    fmt.Println(len(total))
}
