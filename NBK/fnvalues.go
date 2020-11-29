/*
coding: utf-8
Created on 29/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Function calling function
*/

package main

import "fmt"

func compute(fn func(float64, float64) float64, fn2 func(float64, float64) float64) float64 {
	return fn(5, 6) + fn2(3, 4)
}

func main() {

	var add = func(a, b float64) float64 {
		return a + b
	}

	multiplication := func(a, b float64) float64 {
		return a * b
	}
	fmt.Println(add(3, 4))
	fmt.Println(multiplication(3, 4))

	fmt.Println(compute(add, multiplication))

}
