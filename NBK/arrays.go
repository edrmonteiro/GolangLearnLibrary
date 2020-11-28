/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Arrays
*/

package main

import "fmt"

func main() {

	var numbers [7]int
	numbers[0] = 2
	fmt.Println(numbers)

	var numbers1 = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numbers1)

	var numbers2 = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numbers2)

	var numbers3 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers3[5])

	var sum = 0
	for i := 0; i < len(numbers1); i++ {
		fmt.Println(numbers1[i])
		sum += numbers1[i]
	}
	fmt.Println(sum)

}
