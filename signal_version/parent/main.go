package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

// Any code you want.
func StartAll(cmdList []*exec.Cmd) ([]*exec.Cmd, error) {
	for range 20 {
		i, j := rand.IntN(10), rand.IntN(10)
		cmdList[i], cmdList[j] = cmdList[j], cmdList[i]
	}
	for _, cmd := range cmdList {
		err := cmd.Start()
		if err != nil {
			return nil, err
		}
	}
	return cmdList, nil
}


// Any code you want.
func main() {
	// This handle output.txt is empty at start
	os.Remove("output.txt")
	// Any code you want.
	cmdList := []*exec.Cmd {}
	for i := 0 ; i <10 ; i ++ {
		cmd :=  exec.Command("../child/main", fmt.Sprintf("%d", i))
		cmd.Stdout = os.Stdout
		cmdList = append(cmdList, cmd)
		
		// fmt.Println("./auxiliar", fmt.Sprintf("%d", i))
	}
	// Any code you want.
	cmdList, err := StartAll(cmdList)
	if err != nil {
		log.Fatal("Something went wrong:", err)
	}
	// Any code you want.

	// Send SIGUSR1 all child processes
	numero := 0
	// Create a infinite loop
	for {
		// Keep the loop while the programs is searching for the correct child to execute
		for _, cmd := range cmdList {
			Argument, _ := strconv.Atoi(cmd.Args[1])
			// fmt.Println("DEBUG: numeroArgumento =", numeroArgumento, "numero =", numero)
			if numero == Argument {
				// fmt.Println("he entrado", numero)
				cmd.Process.Signal(syscall.SIGUSR1)
				cmd.Wait()
				numero++
			}
			// Apparently this handles a race condition
			time.Sleep(20 * time.Millisecond)
		}
		// The last process increase the number to 10 and breaks the loop
		if numero == 10 {
			break
		}
	}
	



}



