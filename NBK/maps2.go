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

	var profiles map[string]Profile = make(map[string]Profile)

	profiles["Eduardo"] = Profile{
		1.82,
		true,
		"Engineer",
	}
	profiles["Enzo"] = Profile{
		0.60,
		true,
		"Student",
	}

	fmt.Println(profiles)

}
