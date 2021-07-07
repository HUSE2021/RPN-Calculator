package main

import (
	"fmt"
)

func main(){
    var in_string string
	var stack Stack
	var i int

	var answer int

	stack.N = 0
	in_string = "1 3 / 1 +"
	s := strings.Split(in_string, " ")

	for i = 0; i < len(s); i++ {
		if is_num(s[i]) {
			input_num, err := strconv.Atoi(s[i])
			if err != nil {
				fmt.Println("turn int err")
				return
			}
			stack.push(input_num)
		} else {
			var ina int
			var inb int
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
			stack.push(answer)
		}
	}
	answer = stack.GetTop()
	fmt.Println("answer is ", answer)
}


func add(a, b int) int {
	return a + b
}

func sub(a,b int) int {
	return a - b
}

func multiple(a, b int) int {
	return a * b
}

func fraction(a, b int) int {
	return a / b
}

func is_num(a string) bool {
    if a < b
} 


func (s *Stack)push(value int)(){
	for i = s.N-1; i > 0; i-- {
        arry[i] = arry[i-1]
    }
    arry[0] = value
    s.N ++
}


func (s *Stack)pop() (int,int){
    var top int
    var topNext int
    top = arry[s.N--]
    topNext = arry[s.N--]
    return top,topNext
} 


func (s *Stack)printStack(value int)(){
    for i = 0; i < s.N; i++ {
        fmt.Println(s.arry[i])
    }
}


type Stack struct {
	arry [10]int
	N int
}