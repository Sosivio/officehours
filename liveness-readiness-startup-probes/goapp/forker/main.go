package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	messages := make(chan string)
	fmt.Println("Main process")
	for i := 0; i < 10000; i++ {
		go child(messages, i)
		fmt.Println(<-messages)
	}
}

func child(messages chan string, i int) {
	runtime.LockOSThread()
	pid := os.Getpid()

	parentpid := os.Getppid()

	fmt.Printf("The parent process id of %v is %v %v\n", pid, parentpid, i)

	messages <- strconv.Itoa(os.Getpid())

}
