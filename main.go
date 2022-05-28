package main

import (
	"fmt"
	"functional_dequeue/fdq"
	"io"
	"log"
	"os"
)

func main() {
	var q fdq.Dequeue

	if len(os.Args) > 1 {
		q = chooseDequeue(os.Args[1])
	}

REPL:
	for {
		if q == nil {
			q = askImplementation()
			continue
		}

		fmt.Print("> ")

		var operation, data string
		n, err := fmt.Scanf("%s %s\n", &operation, &data)
		if err == io.EOF {
			fmt.Println("EOF on read")
			break
		}
		if n != 1 && err != nil {
			fmt.Printf("Failed to read: %v\n", err)
			break
		}

		var returnData any

		switch operation {
		case "popL":
			returnData, q = q.PopLeft()
			fmt.Printf("pop left: %v\n", returnData)
		case "pushL":
			if n > 1 {
				q = q.PushLeft(data)
			}
		case "popR":
			returnData, q = q.PopRight()
			fmt.Printf("pop right: %v\n", returnData)
		case "pushR":
			if n > 1 {
				q = q.PushRight(data)
			}
		case "print":
			q.Print(os.Stdout)
		case "quit":
			break REPL
		default:
		}
	}
}

func chooseDequeue(implementation string) fdq.Dequeue {
	switch implementation {
	case "dllist":
		return (*fdq.ListQueue)(nil)
	case "twostack":
		return (*fdq.TwoStack)(nil)
	}
	fmt.Printf("unknown dequeue implemention: %q\n", implementation)
	return nil
}

func askImplementation() fdq.Dequeue {
	fmt.Printf("Available implementations: %v\n", fdq.Implementations)
	fmt.Print("Choose dequeue implementation: ")
	var qimp string
	n, err := fmt.Scanf("%s\n", &qimp)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("no implementation named")
	}
	return chooseDequeue(qimp)
}
