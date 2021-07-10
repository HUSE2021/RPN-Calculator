package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Fraction struct {
	n int
	d int
}

type Stack struct {
	arry [20]Fraction
	N    int
}

func main() {
	var in_string string
	var stack Stack
	var i int

	var answer Fraction

	stack.N = 0
	in_string = "4 3 / 1 3 / + 2 * 4 - 2 +"
	s := strings.Split(in_string, " ")

	for i = 0; i < len(s); i++ {
		if is_num(s[i]) {
			var input_fra Fraction
			input_fra = strt2fra(s[i])
			stack.push(input_fra)
		} else {
			var ina Fraction
			var inb Fraction
			ina, inb = stack.pop()
			if s[i] == "+" {
				answer = add(inb, ina)
			} else if s[i] == "-" {
				answer = sub(inb, ina)
			} else if s[i] == "*" {
				answer = multiple(inb, ina)
			} else if s[i] == "/" {
				answer = fraction(inb, ina)
			}
			var f int
			f = gcd(answer.n, answer.d)
			answer.n /= f
			answer.d /= f
			stack.push(answer)
		}
	}
	answer = stack.GetTop()
	if answer.d == 1 {
		fmt.Println("answer is ", answer.n)
	} else if answer.d < 0 {
		fmt.Println("answer is ", -answer.n, "/", -answer.d)
	} else {
		fmt.Println("answer is ", answer.n, "/", answer.d)
	}

}

func (s *Stack) GetTop() int {
	return s.arry[s.N]
}

func (s *Stack) pop() (int, int) {
	var top int
	var topNext int
	top = s.arry[s.N]
	s.N = s.N - 1
	topNext = s.arry[s.N]
	s.N = s.N - 1
	return top, topNext
}

func (s *Stack) push(value int) {
	s.N++
	s.arry[s.N] = value
}

func (s *Stack) printStack() {
	var i int
	for i = 0; i < s.N; i++ {
		fmt.Println(s.arry[i])
	}
}

//Fraction operation
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	var hcf int
	var temp int
	hcf = a
	temp = b

	for hcf != temp {
		if hcf > temp {
			hcf -= temp
		} else {
			temp -= hcf
		}
	}

	return (a * b) / hcf
}

func add(a, b Fraction) Fraction {
	var new_d int
	var c Fraction

	new_d = lcm(a.d, b.d)
	a.n = a.n * new_d / a.d
	b.n = b.n * new_d / b.d

	c.d = new_d
	c.n = a.n + b.n

	return c
}

func sub(a, b Fraction) Fraction {
	//a > b
	var new_d int
	var c Fraction

	new_d = lcm(a.d, b.d)
	a.n = a.n * new_d / a.d
	b.n = b.n * new_d / b.d

	c.d = new_d
	c.n = a.n - b.n

	return c
}

func multiple(a, b Fraction) Fraction {
	var y Fraction
	y.n = a.n * b.n
	y.d = a.d * b.d
	return y
}

func fraction(a, b Fraction) Fraction {
	var y Fraction
	y.n = a.n * b.d
	y.d = a.d * b.n
	return y
}

func is_num(a string) bool {
	return unicode.IsDigit(rune(a[0]))
}

func strt2fra(in string) Fraction {
	var fra Fraction

	input_num, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println("turn int err")
	}

	fra.n = input_num
	fra.d = 1
	return fra
}
