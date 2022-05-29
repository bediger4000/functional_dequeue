package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
	for ; count > 0; count-- {
		switch rand.Intn(4) {
		case 0: // push left
			fmt.Printf("pushL %d\n", count)
		case 1: // pop left
			fmt.Println("popL")
		case 2: // push right
			fmt.Printf("pushR %d\n", count)
		case 3: // pop right
			fmt.Println("popR")
		}
	}
	fmt.Println("ops")
}
