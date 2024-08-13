package pong

import "time"

const (
    Reset      = "\033[0m"
    PongColor  = "\033[36m" 
)

func Pong(c chan<- string, done <-chan bool) {
    for {
        select {
        case c <- PongColor + "pong" + Reset:
            time.Sleep(1 * time.Second)
        case <-done:
            return
        }
    }
}
