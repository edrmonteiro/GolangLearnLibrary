/*
coding: utf-8
Created on 28/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Inference
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	var j = i

	fmt.Printf("Type: %T value %v", j, j)

	var k = 0.64 + 0.3i
	fmt.Printf("Type: %T value %v", k, k)

	const Pi = 3.14

	fmt.Printf("Type: %T value %v", Pi, Pi)

	var number = math.Max(3, 5)

	fmt.Printf("Type: %T value %v", number, number)

}
