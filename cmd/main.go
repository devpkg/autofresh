package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Hi")
		fmt.Println("I am running forever.")
		time.Sleep(10 * time.Second)

		fmt.Fprintln(os.Stderr, "Stderr error message.")

		time.Sleep(10 * time.Second)

	}
}
