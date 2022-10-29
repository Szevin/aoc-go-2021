package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	amount    int
}

type Position struct {
	horizontal int
	vertical   int
	aim        int
}

func (p *Position) addA(c Command) {
	switch c.direction {
	case "forward":
		p.horizontal += c.amount
	case "down":
		p.vertical += c.amount
	case "up":
		p.vertical -= c.amount
	}
}

func (p *Position) addB(c Command) {
	switch c.direction {
	case "forward":
		p.horizontal += c.amount
		p.vertical += c.amount * p.aim
	case "down":
		p.aim += c.amount
	case "up":
		p.aim -= c.amount
	}
}

var commands = []Command{}

func main() {
	f, err := os.Open("2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		direction, _ := strconv.Atoi(input[1])
		commands = append(commands, Command{input[0], direction})
	}

	a()
	b()
}

func a() {
	position := Position{}
	for _, command := range commands {
		position.addA(command)
	}
	fmt.Printf("A: %d\n", position.horizontal*position.vertical)
}

func b() {
	position := Position{}
	for _, command := range commands {
		position.addB(command)
	}
	fmt.Printf("B: %d\n", position.horizontal*position.vertical)
}
