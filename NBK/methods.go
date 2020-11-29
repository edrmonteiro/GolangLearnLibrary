/*
coding: utf-8
Created on 29/11/2020
@author: github.com/edrmonteiro
From: NBK
Language: Go
Title: Methods
*/

package main

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName

}

func FuncFullName(p Person) string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) CapitalizeName() {
	p.FirstName = strings.ToUpper(p.FirstName)
}

func main() {
	someone := Person{"José", "de Alencar"}
	fmt.Println(someone)
	fmt.Println(someone.FirstName)

	fmt.Println(someone.FullName())
	fmt.Println(FuncFullName(someone))

	someone.CapitalizeName()
	fmt.Println(FuncFullName(someone))

	someone1 := Person{"José", "de Alencar"}
	p := &someone1
	fmt.Println(p.FullName())

}
