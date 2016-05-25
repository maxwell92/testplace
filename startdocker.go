package main
import  (
    "fmt"
    "os"
    "strconv"
    "os/exec"
)

func main() {
    fmt.Println("start docker for 1")    
    cmd := exec.Command("./ScaleVmUp", strconv.Itoa(n))
    str, err := cmd.Output()
    if err != nil {
        log.Println(err)
    }
    fmt.Println("started")    
    os.Exit(0) 
}
