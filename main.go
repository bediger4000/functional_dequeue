package main

import (
	"fmt"
	"functional_dequeue/fdq"
	"os"
)

func main() {
	var q fdq.Dequeue

	switch os.Args[1] {
	case "dllist":
		q = fdq.NewDllist()
	}

	var data any

	for idx := 2; idx < len(os.Args); {

		thing := os.Args[idx]

		switch thing {
		case "popL":
			idx++
			data, q = q.PopLeft()
			fmt.Printf("pop left: %v\n", data)
		case "pushL":
			q = q.PushLeft(os.Args[idx+1])
			idx += 2
		}
	}
}
