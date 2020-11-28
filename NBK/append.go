/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Append
*/

package main

import "fmt"

func main() {
	var names []string
	var names1 = append(names, "John")
	names1 = append(names1, "John", "Ane", "Roger")

	fmt.Println(names)
	fmt.Println(names1)

	fmt.Println(names1)
	fmt.Println("len= %d; cap = %d", len(names1), cap(names1))

}
