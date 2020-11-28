/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: function
*/

package main

import "fmt"

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(power(2))
	fmt.Println(power1(2))
}

func power(a int) (int, int) {
	var square int = a * a
	var cubic int = a * a * a
	return square, cubic
}

func power1(a int) (square int, cubic int) {
	square = a * a
	cubic = a * a * a
	return
}

func add(x int, y int) int {
	return x + y
}
