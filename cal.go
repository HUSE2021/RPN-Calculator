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
