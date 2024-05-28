package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello from goroutine")
	fmt.Println("hello from main")
	time.Sleep(100 * time.Millisecond) // Pausa breve para dar tiempo a la goroutine a ejecutarse
}
