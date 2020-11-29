/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Pointers
*/

package main

import "fmt"

func main() {
	a := 32
	p := &a

	fmt.Println(a)
	*p = 18
	*p++

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(a)

	var p1 *int = &a
	fmt.Println(*p1)

}
