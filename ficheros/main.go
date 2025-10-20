package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
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

// This handle the removal of lock files
func borrarLock() {
	for i := 0; i <= 10; i++ {
		err := os.Remove(".lock" + fmt.Sprintf("%d", i))
		if err != nil {
			log.Fatal("Something went wrong:", err)
		}
	}
}


// Any code you want.
func main() {
	// This handle output.txt is empty at start
	os.Remove("output.txt")
	// Any code you want.
	cmdList := []*exec.Cmd {}
	for i := 0 ; i <10 ; i ++ {
		cmd :=  exec.Command("./auxiliar", fmt.Sprintf("%d", i))
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

	// Create first lock 0
	os.WriteFile(".lock0", []byte(""), 0644)
	// Wait for all commands to finish
	for _, cmd := range cmdList {
		err := cmd.Wait()
		
		if err != nil {
			log.Fatal("Something went wrong:", err)
		}
	}
	borrarLock()


}