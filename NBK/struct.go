/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Struct
*/

package main

import "fmt"

func main() {

	type Position struct {
		x int
		y int
	}
	pos := Position{40, 67}
	pos.y = 51
	fmt.Println(pos)

	var pos1 Position = Position{y: 47}
	fmt.Println(pos1)

}
