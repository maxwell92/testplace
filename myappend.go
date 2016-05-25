package main
import (
    "log"
    "os"
)

func main() {
    fp, err := os.OpenFile("./a", os.O_RDWR | os.O_APPEND, 0644)
    if err != nil {
	log.Println(err)
    }

    fp.WriteString("Hello SHeep") 
}


