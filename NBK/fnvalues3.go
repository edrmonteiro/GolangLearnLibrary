/*
coding: utf-8
Created on 29/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Function calling function - closure
*/

package main

import "fmt"

func aggregator() func(a int) int {
	var sum int = 0
	//closure
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	var add1 = aggregator()
	add1(2)
	fmt.Println(add1(21))
	fmt.Println(add1(13))
	fmt.Println("------")

	var add2 = aggregator()
	fmt.Println(add2(21))
	fmt.Println(add2(13))
}
