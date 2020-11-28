package main

import "fmt"

func main() {
	var age, height int = 35, 188
	var name = "Eduardo"

	var (
		age1    int = 5
		height1     = 80
		name1       = "Enzo"
		name2   string
	)

	age2 := 36
	name3, height3, age3 := "Lara", 120, 9

	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//boll true /false
	//complex64 complex128 a + bi
	//byte, rune

	fmt.Println(age)
	fmt.Println("My name is", name, "I'm", age, "and", height, "tall")
	fmt.Println("My name is", name1, "I'm", age1, "and", height1, "tall")
	fmt.Println("My name is", name2, "I'm", age2, "and", height1, "tall")
	fmt.Println("My name is", name3, "I'm", age3, "and", height3, "tall")

	var height4 = 1.86
	fmt.Printf("Tipo: %T Valor: %v\n", height4, height4)

}
