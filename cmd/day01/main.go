package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := "inputs/day01.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("failed to open file %s: %w", filename, err))
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(fmt.Errorf("failed to reset file for Part 2: %w", err))
	}

	fmt.Printf("Part 2: %d\n", part2(f))

}

func part1(i io.Reader) int {
	d := &dial{
		pos: 50,
	}

	s := bufio.NewScanner(i)

	for s.Scan() {
		d.turn(s.Text())
	}

	return d.acc
}

func part2(i io.Reader) int {
	d := &improvedDial{
		pos: 50,
	}

	s := bufio.NewScanner(i)

	for s.Scan() {
		d.turn(s.Text())
	}

	return d.acc
}

type dial struct {
	pos int
	acc int
}

func (d *dial) turn(instruction string) {
	direction := 1
	switch instruction[0] {
	case 'R':
		direction = 1
	case 'L':
		direction = -1
	}

	magnitude, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(fmt.Errorf("unexpected magnitude '%s': %w", instruction[1:], err))
	}

	magnitude = magnitude % 100
	d.pos = d.pos + (magnitude * direction)
	if d.pos < 0 {
		d.pos = d.pos + 100
	}

	d.pos = d.pos % 100

	if d.pos == 0 {
		d.acc++
	}
}

type improvedDial struct {
	pos int
	acc int
}

func (d *improvedDial) turn(instruction string) {
	direction := 1
	switch instruction[0] {
	case 'R':
		direction = 1
	case 'L':
		direction = -1
	}

	magnitude, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(fmt.Errorf("unexpected magnitude '%s': %w", instruction[1:], err))
	}

	d.acc += magnitude / 100
	magnitude = magnitude % 100

	if magnitude == 0 {
		return
	}

	start := d.pos
	d.pos = d.pos + (magnitude * direction)
	if d.pos < 0 {
		d.pos = 100 + d.pos
	} else if d.pos >= 100 {
		d.pos = d.pos % 100
	}
	
	if start == 0 {
		return
	}

	if d.pos == 0 {
		d.acc++
	} else if direction == -1 && start < d.pos {
		d.acc++
	} else if direction == 1 && start > d.pos {
		d.acc++
	}

}
