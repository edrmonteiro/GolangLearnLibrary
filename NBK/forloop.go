/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: forloop
*/

package main

import "fmt"

func main() {
	var add int = 0

	for i := 1; i <= 10; i += 2 {
		add = add + i
	}
	fmt.Println(add)
}
