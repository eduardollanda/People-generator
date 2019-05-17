package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(genereteName())
}

func genereteName() (string, string) {

	firstNameMan := [...]string{
		"Eduardo",
		"Cayo",
		"Gustavo",
		"Igor",
		"Matheus",
		"Mateus",
	}
	firstNamewoman := [...]string{
		"Lígia",
		"Ellen",
		"karol",
		"Joaquina",
		"Fernanda",
		"Ana paula",
	}

	return firstNameMan[rand.Intn(len(firstNameMan))], firstNamewoman[rand.Intn(len(firstNamewoman))]
}
