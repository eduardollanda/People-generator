package main

import (
	"fmt"
	"math/rand" // Para gerar numeros aleatorios
	"net/http"
	"strconv"
	"text/template"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/mulher", mulherpage)
	http.HandleFunc("/homem", homempage)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	var option int64
	tpl, _ := template.ParseFiles("Front/index.html")
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, option)
}
func homempage(w http.ResponseWriter, r *http.Request) {

	homem, _ := genereteName()
	cpf := gerarCpf()
	rg := gerarRG()

	tpl, _ := template.ParseFiles("Front/gerador.html")
	data := map[string]string{
		"Nome": homem,
		"CPF":  cpf,
		"RG":   rg,
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)

}

func mulherpage(w http.ResponseWriter, r *http.Request) {

	_, mulher := genereteName()
	cpf := gerarCpf()
	rg := gerarRG()
	tpl, _ := template.ParseFiles("Front/gerador.html")
	data := map[string]string{
		"Nome": mulher,
		"CPF":  cpf,
		"RG":   rg,
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)

}

func genereteName() (homem, mulher string) {

	firstNameMan := [...]string{
		"Cayo",
		"Eduardo",
		"Geriston",
		"Gustavo",
		"Igor",
		"Mateus",
		"Matheus",
	}
	firstNamewoman := [...]string{
		"Ana paula",
		"Ellen",
		"Fernanda",
		"Joaquina",
		"Lígia",
		"karol",
	}
	lastName := [...]string{
		"Costa",
		"Ferreira",
		"Lima",
		"Melo",
		"Oliveira",
		"Pereira",
		"Rodrigues",
		"Santos",
		"Silva",
		"Sousa",
		"Souza",			
	}
	completeNameMan := (firstNameMan[rand.Intn(len(firstNameMan))] + " " + lastName[rand.Intn(len(lastName))])
	completeNamewoman := (firstNamewoman[rand.Intn(len(firstNamewoman))] + " " + lastName[rand.Intn(len(lastName))])

	return completeNameMan, completeNamewoman
}

func gerarCpf() string {

	var (
		cpf            []int
		primeiraSoma   int
		primeiroDigito int
		segundaSoma    int
		segundoDigito  int
		cpfToString    string
	)

	//criando os 9 primeiros digitos do cpf
	for cont1 := 0; cont1 < 9; cont1++ {
		aleatorio := rand.Intn(10)
		cpf = append(cpf, aleatorio)
	}

	// Calculo para validar os dois ultimos digitos do CPF
	// Calculo retirado de https://www.geradorcpf.com/algoritmo_do_cpf.htm
	primeiraSoma = (cpf[0] * 10) + (cpf[1] * 9) + (cpf[2] * 8) + (cpf[3] * 7) + (cpf[4] * 6) + (cpf[5] * 5) + (cpf[6] * 4) + (cpf[7] * 3) + (cpf[8] * 2)

	if (primeiraSoma % 11) < 2 {
		primeiroDigito = 0
	} else {
		primeiroDigito = 11 - (primeiraSoma % 11)
	}
	cpf = append(cpf, primeiroDigito)

	segundaSoma = (cpf[0] * 11) + (cpf[1] * 10) + (cpf[2] * 9) + (cpf[3] * 8) + (cpf[4] * 7) + (cpf[5] * 6) + (cpf[6] * 5) + (cpf[7] * 4) + (cpf[8] * 3) + (cpf[9] * 2)

	if (segundaSoma % 11) < 2 {
		segundoDigito = 0
	} else {
		segundoDigito = 11 - (segundaSoma % 11)
	}

	cpf = append(cpf, segundoDigito)

	for elemento := range cpf {
		elementoToString := strconv.Itoa(cpf[elemento])
		cpfToString = (cpfToString + elementoToString)
	}
	return cpfToString

}

func gerarRG() string {
	var (
		rg             []int
		primeiraSoma   int
		primeiroDigito int
		subtraendo     int
		rgToString     string
	)

	for cont1 := 0; cont1 < 8; cont1++ {
		aleatorio := rand.Intn(10)
		rg = append(rg, aleatorio)
	}
	// calculo retirado de https://www.ngmatematica.com/2014/02/como-determinar-o-digito-verificador-do.html
	primeiraSoma = (rg[0] * 2) + (rg[1] * 3) + (rg[2] * 4) + (rg[3] * 5) + (rg[4] * 6) + (rg[5] * 7) + (rg[6] * 8) + (rg[7] * 9)

	subtraendo = (primeiraSoma % 11)

	primeiroDigito = (11 - subtraendo)

	if primeiroDigito == 11 {
		for elemento := range rg {
			elementoToString := strconv.Itoa(rg[elemento])
			rgToString = (rgToString + elementoToString)
		}
		primeiroDigito = 0
		rgToString = (rgToString + strconv.Itoa(primeiroDigito))
	} else if primeiroDigito == 10 {
		for elemento := range rg {
			elementoToString := strconv.Itoa(rg[elemento])
			rgToString = (rgToString + elementoToString)
		}
		rgToString = (rgToString + "x")
	} else {
		for elemento := range rg {
			elementoToString := strconv.Itoa(rg[elemento])
			rgToString = (rgToString + elementoToString)
		}
		rgToString = (rgToString + strconv.Itoa(primeiroDigito))
	}
	return rgToString
}
