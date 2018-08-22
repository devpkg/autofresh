package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("New Message")
	for {
		fmt.Println("I am running forever.")
		time.Sleep(1 * time.Second)

		fmt.Fprintln(os.Stderr, "Stderr error message.")

		time.Sleep(1 * time.Second)

	}
}
