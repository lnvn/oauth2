package main

import "fmt"

type ByteSize int

const (
	_ = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Printf("%d, %b\n", KB, KB)
	fmt.Printf("%d, %b\n", MB, MB)
	fmt.Printf("%d, %b\n", GB, GB)
}