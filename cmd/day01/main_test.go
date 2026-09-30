package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	r := strings.NewReader(input)
	assert.Equal(t, 3, part1(r))
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	r := strings.NewReader(input)
	assert.Equal(t, 6, part2(r))
}

func TestImprovedDial(t *testing.T) {
	t.Run("example from problem statement", func(t *testing.T) {
		d := &improvedDial{pos: 50, acc: 0}
		d.turn("R1000")

		assert.Equal(t, 10, d.acc)
		assert.Equal(t, 50, d.pos)
	})

	t.Run("landing on 0 while turning left", func(t *testing.T) {
		d := &improvedDial{pos:50, acc: 0}
		d.turn("L50")

		assert.Equal(t, 0, d.pos)
		assert.Equal(t, 1, d.acc)
	})

	t.Run("leaving zero while turning left", func(t *testing.T) {
		d := &improvedDial{pos: 0, acc: 0}
		d.turn("L5")

		assert.Equal(t, 95, d.pos)
		assert.Equal(t, 0, d.acc)
	})

	t.Run("crossing zero while turning right", func(t *testing.T) {
		d := &improvedDial{pos: 95, acc: 0}
		d.turn("R6")

		assert.Equal(t, 1, d.pos)
		assert.Equal(t, 1, d.acc)
	})
}
