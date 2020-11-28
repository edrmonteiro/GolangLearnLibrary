/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Logic
*/

package main

import "fmt"

func main() {
	var x int = 1

	fmt.Println("x > 3 && x <7", x > 3 && x < 7)
	fmt.Println("x > 3 && x <7", x > 3 || x < 7)

}
