package main

import (
	"fmt"
	"os"
	"strconv"
)

// Theoretically, this function write the line
func main() {
	for {
		_, err := os.Stat(".lock" + os.Args[1])
		if err == nil {
			break
		}
	} 
	numero := os.Args[1]
	fmt.Println(numero)
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(numero + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	// Create a lock file
	lock, _:= strconv.Atoi(numero)
	lock++
	os.WriteFile(".lock"+strconv.Itoa(lock), []byte(""), 0644)
}