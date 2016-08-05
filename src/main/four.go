package main

import (
	"log"
	"strings"
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		data, _, _ := reader.ReadLine()
		opString := string(data)

		if opString == "stop" {
			running = false
		}

		log.Println(opString)
		fmt.Print(opString)

		opString = strings.Replace(opString, "+", ",+,", -1)
		opString = strings.Replace(opString, "-", ",-,", -1)
		opString = strings.Replace(opString, "*", ",*,", -1)
		opString = strings.Replace(opString, "/", ",/,", -1)
		opString = strings.Replace(opString, "(", ",(,", -1)
		opString = strings.Replace(opString, ")", ",),", -1)
		opString = strings.Replace(opString, ",,", ",", -1)
		opArr := strings.Split(opString, ",")
		middle_slice := []string{}
		operator_stack := stack.NewStack();

		for _, value := range opArr {
			if isOp(value) {
				for true {
					top := operator_stack.Peak()
					if (top == nil || value == "(") {
						operator_stack.Push(value)
						break
					}
					v, _ := top.(string)
					if prior(value) > prior(v) {
						operator_stack.Push(value)
						break
					} else {
						op := operator_stack.Pop()
						v, _ := op.(string)
						if v == "(" {
							break
						}
						middle_slice = append(middle_slice, v)
					}
				}
			} else {
				if value != "" {
					middle_slice = append(middle_slice, value)
				}
			}
		}

		for !operator_stack.Empty() {
			op := operator_stack.Pop()
			v, _ := op.(string)
			middle_slice = append(middle_slice, v)
		}

		//log.Println(middle_slice)

		for _, value := range middle_slice {
			if !isOp(value) {
				num, _ := strconv.Atoi(value)
				operator_stack.Push(num)
			} else {
				op1 := operator_stack.Pop()
				num1, _ := op1.(int)

				op2 := operator_stack.Pop()
				num2, _ := op2.(int)

				switch value {
				case "+":
					operator_stack.Push(num2 + num1)
				case "-"://-
					operator_stack.Push(num2 - num1)
				case "*"://*
					operator_stack.Push(num2 * num1)
				case "/":// /
					operator_stack.Push(num2 / num1)
				}
			}
		}
		for !operator_stack.Empty() {
			op := operator_stack.Pop()
			v, _ := op.(int)
			fmt.Println("=", v)
		}
	}
}

func isOp(op string) bool {
	switch op {
	case "+":
		return true;
	case "-"://-
		return true;
	case "*"://*
		return true;
	case "/":// /
		return true;
	case "(":
		return true;
	case ")":
		return true;
	default:
		return false
	}
}

func prior(op string) int {
	switch op {
	case "+":
		return 3;
	case "-"://-
		return 3;
	case "*"://*
		return 4;
	case "/":// /
		return 4;
	default:
		return -1
	}
}