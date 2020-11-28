/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: while
*/

package main

import "fmt"

func main() {
	var add int = 0

	i := 1
	for i < 10 {
		add = add + i
		i++
	}

	fmt.Println(add)

	i = 1
	for {
		add = add + i
		i++
		if i > 5000 {
			break
		}
	}
	fmt.Println(i)

}
