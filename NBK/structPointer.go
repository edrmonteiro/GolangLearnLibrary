/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Struct Pointers
*/

package main

import "fmt"

func main() {

	type Position struct {
		x int
		y int
	}
	var pos Position = Position{47, 81}
	var k *Position = &pos
	fmt.Println(k)

	k.y = 33

	fmt.Println(pos.y)

}
