package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSimpleCapcha(t *testing.T) {
	solver := new()

	solution := solver.checkCapcha("1122")

	if solution != 3 {
		t.Error("Incorrect answer")
	}
}

func TestCheckRepeatativeCapcha(t *testing.T) {
	solver := new()

	solution := solver.checkCapcha("1111")

	if solution != 4 {
		t.Error("Incorrect answer")
	}
}

func TestCheckCapchaWithZeroAnswer(t *testing.T) {
	solver := new()

	solution := solver.checkCapcha("1234")

	if solution != 0 {
		t.Error("Incorrect answer")
	}
}

func TestCheckCapchaWithCycle(t *testing.T) {
	solver := new()

	solution := solver.checkCapcha("91212129")
	expectedSolution := 9
	assert.Equal(t, expectedSolution, solution)
}
