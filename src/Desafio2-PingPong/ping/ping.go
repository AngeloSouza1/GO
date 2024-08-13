package ping

import "time"

const (
    Reset      = "\033[0m"
    PingColor  = "\033[33m"  
)

func Ping(c chan<- string, done <-chan bool) {
    for {
        select {
        case c <- PingColor + "ping" + Reset:
            time.Sleep(1 * time.Second)
        case <-done:
            return
        }
    }
}
