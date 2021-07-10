package main

import (
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	expected := "4/3"
	result:= answerType(calculate("4 3 / 1 3 / + 2 * 4 - 2 +"))
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}
func Test2(t *testing.T) {
	expected := "5"
	result:= answerType(calculate("1 1 1 1 1 1 + + + +"))
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}
func Test3(t *testing.T) {
	expected := "1/2"
	result:= answerType(calculate("1 3 / 1 6 / +"))
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}

func Test5(t *testing.T) {
	expected := "1"
	result:= answerType(calculate("1 2 / 1 2 / +"))
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}


func answerType(answer Fraction) string {
	if answer.d == 1 {
		return strconv.Itoa(answer.n)
	} else if answer.d < 0 {
		return strconv.Itoa(-answer.n)+ "/"+strconv.Itoa(-answer.d)
	} else {
		return strconv.Itoa(answer.n)+ "/"+strconv.Itoa(answer.d)
	}
}