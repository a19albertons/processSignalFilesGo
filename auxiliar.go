package main

import (
	"fmt"
	"os"
)

// Theoretically, this function write the line
func main() {
	numero := os.Args[1]
	println("holiwy")
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
}