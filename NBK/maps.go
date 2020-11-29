/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Maps
*/

package main

import "fmt"

func main() {
	var weights map[string]int

	weights = make(map[string]int)

	weights["Ane"] = 60
	weights["Mary"] = 65

	fmt.Println(weights)
	fmt.Println(weights["Ane"])
	fmt.Println(weights["John"])

	value, exists := weights["John"]
	if exists {
		fmt.Println(value)
		fmt.Println(exists)
	}

}
