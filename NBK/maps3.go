/*
coding: utf-8
Created on 29/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Maps with structures
*/

package main

import "fmt"

type Profile struct {
	Height     float64
	Active     bool
	Occupation string
}

func main() {
	var profiles map[string]Profile = map[string]Profile{
		"Eduardo": Profile{1.82, true, "Engineer"},
		"Enzo":    Profile{0.60, false, "Student"},
	}
	fmt.Println(profiles)
}
