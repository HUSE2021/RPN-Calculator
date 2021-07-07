package main

import (
	"fmt"
)

func main(){
    var a string
    
    var answer int
    
	fmt.Scanln(&a)

    if(s == "+"){
        answer = add(a, b)
    }else if(s == "-"){
        answer = sub(a, b)
    }else if(s == "*"){
        answer = multiple(a, b)
    }else if(s == "/"){
        answer = fraction(a, b)
    }
        
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
    if a <
} 


func (s *Stack)push(value int)(){
	for i = s.N-1; i > 0; i-- {
        arry[i] = arry[i-1]
    }
    arry[0] = value
    s.N ++
}

func (s *Stack)printStack(value int)(){
    for i = 0; i < s.N; i++ {
        fmt.Println(s.arry[i])
    }
}
