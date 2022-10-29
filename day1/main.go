package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Depths []int

func (d *Depths) append(value int) {
	*d = append(*d, value)
}

func (d Depths) len() int {
	return len(d)
}

var depths = Depths{}

func main() {
	f, _ := os.Open("1.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		depths.append(value)
	}

	a()
	b()
}

func a() {
	increased := 0
	for i := 1; i < depths.len(); i++ {
		if depths[i] > depths[i-1] {
			increased++
		}
	}

	fmt.Printf("A: %d\n", increased)
}

func b() {
	increased := 0
	for i := 3; i < depths.len(); i++ {
		if depths[i] > depths[i-3] {
			increased++
		}
	}

	fmt.Printf("B: %d\n", increased)
}
