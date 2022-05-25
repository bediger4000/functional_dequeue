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

	fmt.Printf("%T %+v\n", q, q)
}
