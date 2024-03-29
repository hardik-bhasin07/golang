package main

import "fmt"

const (
	c0 = iota
	c1
	c2
)

func add(x int, y int) int {
	return x + y
}

type ByteSize int

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t\t %b\n", TB, TB)
	fmt.Printf("%d \t\t\t %b\n", PB, PB)
	fmt.Printf("%d \t\t\t %b\n", EB, EB)
}
