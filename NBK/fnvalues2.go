/*
coding: utf-8
Created on 29/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Function calling function
*/

package main

import "fmt"

func greeting(name string) func() string {
	//return "Hello, " + name + "!"
	//return fmt.Sprintf("Hello, %s!", name)
	return func() string {
		return fmt.Sprintf("Hello, %s!", name)
	}
}

func main() {
	//fmt.Println(greeting("Eduardo"))

	greetingEduardo := greeting("Eduardo")
	fmt.Println(greetingEduardo())

	greetingEnzo := greeting("Enzo")
	fmt.Println(greetingEnzo())
}
