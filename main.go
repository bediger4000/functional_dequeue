package main

import (
	"flag"
	"fmt"
	"functional_dequeue/fdq"
	"io"
	"log"
	"os"
)

func main() {
	implementation := flag.String("i", "twostack", "choice of implementation")
	quiet := flag.Bool("q", false, "non-interactive output")
	flag.Parse()

	var q fdq.Dequeue

	if flag.NArg() > 0 {
		q = chooseDequeue(flag.Arg(0))
	} else {
		q = chooseDequeue(*implementation)
	}

REPL:
	for {
		if q == nil {
			q = askImplementation()
			continue
		}

		if !*quiet {
			fmt.Print("> ")
		}

		var operation, data string
		n, err := fmt.Scanf("%s %s\n", &operation, &data)
		if err == io.EOF {
			// fmt.Println("EOF on read")
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
			if !*quiet {
				fmt.Print("pop left: ")
			}
			fmt.Printf("%v\n", returnData)
		case "pushL":
			if n > 1 {
				q = q.PushLeft(data)
			}
		case "popR":
			returnData, q = q.PopRight()
			if !*quiet {
				fmt.Print("pop right: ")
			}
			fmt.Printf("%v\n", returnData)
		case "pushR":
			if n > 1 {
				q = q.PushRight(data)
			}
		case "print":
			q.Print(os.Stdout)
		case "quit":
			break REPL
		case "type":
			fmt.Printf("%s\n", q.Type())
		case "ops":
			dqOps, internalOps := q.Operations()
			fmt.Printf("%d/%d => %.3f\n", dqOps, internalOps, float64(internalOps)/float64(dqOps))
		case "new":
			if n > 1 {
				q = chooseDequeue(data)
			} else {
				q = askImplementation()
			}
		default:
		}
	}
}

func chooseDequeue(implementation string) fdq.Dequeue {
	if dq, ok := fdq.NewFunctions[implementation]; ok {
		return dq
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
