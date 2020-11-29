/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Defer
*/

package main

import "fmt"

func print() string {
	fmt.Println("Printing...")
	return "value to print"
}

func main() {

	//defer FILO queue
	defer fmt.Println(print())
	defer fmt.Println("second")
	fmt.Println("third")

}
