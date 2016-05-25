package main
import "log"

func main() {
    n = 1 
    cmd := exec.Command("./ScaleVmUp", strconv.Itoa(n))
	if _, err := cmd.Output(); err != nil {
	    log.Println(err)
	}

}
