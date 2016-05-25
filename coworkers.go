package main
import (
    "log"
    "sync"
    "os"
    //"bufio"
    "strconv"
)

var fp *os.File

func write(id int, wg *sync.WaitGroup) {
    log.Println("This is #" + strconv.Itoa(id) + " worker!")
    for {
        if id == 0 {
            fp.WriteString("Hello Mushroom")
        }
        if id == 1 {
            fp.WriteString("Hello Sheep")
        }
        if id == 2 {
            fp.WriteString("Hello Ugly")
        }
    }
    wg.Done()
}

func main() {
    num := 3
    var wg sync.WaitGroup
    fp, _  = os.OpenFile("./bbb", os.O_RDWR | os.O_APPEND, 0644)
    /*if err != nil {
        log.Println(err)
    }*/
    wg.Add(3)
    for i := 0; i < num; i++ {
        go write(i, &wg)
    }
    wg.Wait()
}
