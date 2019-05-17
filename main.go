package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var sexoEscoliho int
	fmt.Println("Homem ou mulher ? \n homem = 1\n mulher = 2")
	fmt.Scan(&sexoEscoliho)

	homem, mulher := genereteName()

	if sexoEscoliho == 1 {
		fmt.Println(homem)

	} else if sexoEscoliho == 2 {
		fmt.Println(mulher)
	}

}

func genereteName() (homem, mulher string) {

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
	lastName := [...]string{
		"Silva",
		"Santos",
		"Oliveira",
		"Souza",
		"Lima",
		"Pereira",
		"Ferreira",
		"Costa",
		"Rodrigues",
	}
	completeNameMan := (firstNameMan[rand.Intn(len(firstNameMan))] + " " + lastName[rand.Intn(len(lastName))])
	completeNamewoman := (firstNamewoman[rand.Intn(len(firstNamewoman))] + " " + lastName[rand.Intn(len(lastName))])

	return completeNameMan, completeNamewoman
}
