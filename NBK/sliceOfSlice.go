/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Slice of Slice
*/

package main

import "fmt"

func main() {
	table := [][]string{
		[]string{"x", "-", "-"},
		[]string{"-", "x", "-"},
		[]string{"-", "-", "x"},
	}
	fmt.Println(table)
	fmt.Println(table[1])

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Println(table[j][i])
		}
	}
}
