package main
import "fmt"

func main() {
	data := make([]float32, 10, 10)
	fmt.Println(len(data))
	data[0] = 1.0
	data[1] = 2.0
	data[2] = 3.0
	fmt.Println(len(data))
}
