package main

import "fmt"

var (
	Version = "unknown"
	Build   = "unknown"
)

func main() {
	fmt.Println("go-realworld-clean forked by kazufusa")
	fmt.Printf("version: %s\n", Version)
	fmt.Printf("build: %s\n", Build)
}
