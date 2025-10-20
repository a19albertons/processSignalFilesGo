package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1)
	for {
		sig := <-sigChan
		if sig == syscall.SIGUSR1 {
			// Read the argument number
			numero := os.Args[1]
			// Handle the file writing
			f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer f.Close()
			f.Write([]byte(numero + "\n"))
			f.Close()
			// Write the number to terminal
			fmt.Println(numero)
			// Apparently this handles a race condition
			time.Sleep(10 * time.Millisecond)
			break
		}
	}
	
	
}