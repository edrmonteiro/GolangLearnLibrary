/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Arithmetic operations
*/

package main

import "fmt"

func main() {
	var a int = 23
	var b int = 7

	fmt.Println("a + b=", a+b)
	fmt.Println("a - b=", a-b)
	fmt.Println("a * b=", a*b)
	fmt.Println("a / b=", a/b)
	fmt.Println("a % b=", a%b)

	var c = 23.0236
	var d int = 7

	fmt.Println("c + d=", c+float64(d))
	fmt.Println("c - d=", c-float64(d))
	fmt.Println("c * d=", c*float64(d))
	fmt.Println("c / d=", c/float64(d))
}
