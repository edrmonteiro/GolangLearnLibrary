/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Range
*/

package main

import "fmt"

func main() {

	numbers := []int{10, 20, 300, 40}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for i, value := range numbers {
		fmt.Println(i)
		fmt.Println(value)
	}

	for _, value := range numbers {
		fmt.Println(value)
	}

}
