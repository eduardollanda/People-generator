package main

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"          // Mostra mensagens no console
	"math/rand"
	"net/http" // Gerencia URLs e Servidor Web
	"strconv"
	"text/template" // Gerencia templates

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

func main() {
	//Mostra no log do terminal
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/showMan", showRndMan)
	http.HandleFunc("/showWoman", showRndWoman)
	http.HandleFunc("/", index)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
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

type NamesMan struct {
	IdM       int
	FirstName string
	LastName  string
}

type NamesWoman struct {
	IdW       int
	FirstName string
	LastName  string
}

func dbConnMen() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345"
	dbName := "plpInfo"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//var tmpl = template.Must(template.ParseFiles("tt.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"Title": "Gerador :)",
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func showRndMan(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConnMen()

	// Realiza a consulta com banco de dados e trata erros, pega apenas um com valor aleatorio
	selDB, err := db.Query("select FirstName from NamesMan order by rand() limit 1") // Query para pegar FirstName Aleatorio no banco NamesMan
	selDB1, err := db.Query("select LastName from NamesMan order by rand() limit 1") // Query para pegar LastName Aleatorio no banco NamesMan
	//trata erros
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := NamesMan{}

	// Realiza a estrutura de repetição pegando Firstname no banco
	for selDB.Next() {
		var firstName string

		err = selDB.Scan(&firstName)

		n.FirstName = firstName
	}

	// Realiza a estrutura de repetição pegando Lastname no banco
	for selDB1.Next() {
		var lastName string

		err = selDB1.Scan(&lastName)

		n.LastName = lastName
	}

	//Seta e executa template
	tpl, _ := template.ParseFiles("gerador.html")
	cpf := gerarCpf()
	data := map[string]string{
		"Nome":      n.FirstName,
		"SobreNome": n.LastName,
		"Cpf":       cpf,
	}

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)

	//Fecha conexão com o banco
	defer db.Close()
}

func showRndWoman(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConnMen()

	// Realiza a consulta com banco de dados e trata erros, pega apenas um com valor aleatorio
	selDB, err := db.Query("select FirstName from NamesWoman order by rand() limit 1") // Query para pegar FirstName Aleatorio no banco NamesWoman
	selDB1, err := db.Query("select LastName from NamesWoman order by rand() limit 1") // Query para pegar FirstName Aleatorio no banco NamesWoman

	//trata erros
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := NamesWoman{}

	// Realiza a estrutura de repetição pegando Firstname no banco
	for selDB.Next() {
		var firstName string

		err = selDB.Scan(&firstName)

		n.FirstName = firstName
	}

	// Realiza a estrutura de repetição pegando Lastname no banco
	for selDB1.Next() {
		var lastName string

		err = selDB1.Scan(&lastName)

		n.LastName = lastName
	}

	//Seta e executa template
	tpl, _ := template.ParseFiles("gerador.html")
	cpf := gerarCpf()
	data := map[string]string{
		"Nome":      n.FirstName,
		"SobreNome": n.LastName,
		"Cpf":       cpf,
	}

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)

	//Fecha conexão com o banco
	defer db.Close()
}
