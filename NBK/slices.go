/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Slices
*/

package main

import "fmt"

func main() {

	var numbers1 = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numbers1[2:4])

	fmt.Println(numbers1[2:])

	fmt.Println(numbers1[:5])

	var names = [3]string{"Ane", "Joseph", "mary"}
	var p1 []string = names[0:2]
	p1[0] = "Hamilton"
	fmt.Println(p1)
	fmt.Println(names)

	var names1 = []string{"Ane", "Joseph", "mary"}
	fmt.Println("%T", names1)
	fmt.Println("len=", len(names1))

	fmt.Println("cap=", cap(names1))

	var names2 = []string{"Ane", "Joseph", "mary"}
	fmt.Println(names2)
	fmt.Println(names2[:2])
	fmt.Println(len(names2[:2]))
	fmt.Println(cap(names2[:2]))

	fmt.Println(names2[1:])
	fmt.Println(len(names2[1:]))
	fmt.Println(cap(names2[1:]))

}
