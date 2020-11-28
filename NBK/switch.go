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
	var value int = 8

	switch {
	case value > 9:
		fmt.Println("excelente")
	case value > 7:
		fmt.Println("Very good")
	default:
		fmt.Println("Not so good")
	}
}
