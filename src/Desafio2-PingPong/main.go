package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"Desafio2-PingPong/ping"  
	"Desafio2-PingPong/pong"  
)
func main() {
	c := make(chan string)
	done := make(chan bool)

	go ping.Ping(c, done)
	go pong.Pong(c, done)

	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		done <- true
	}()

loop:
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
			time.Sleep(1 * time.Second)
		case <-done:
			fmt.Println("Stopping...")
			close(c)
			break loop
		}
	}
}
