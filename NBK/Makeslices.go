/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Make Slices
*/

package main

import "fmt"

func main() {

	var numbers = make([]int, 5)
	fmt.Println(numbers)

	var numbers1 = make([]int, 0, 5)
	fmt.Println(numbers1[0:2])
	fmt.Println(numbers1[0:5])

}
