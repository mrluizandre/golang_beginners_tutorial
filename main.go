package main

import (
	"fmt"
	"strings"
)

func main() {

	// DECLARAÇÃO DE VARIÁVEIS
	// Go infere o tipo de dado quando definimos o valor na declaração
	var conferenceName = "Conferência Go" // alternativa => conferenceName := 50, e não funciona pra declarar o tipo explicitamente
	const referenceTickets = 50           // não há alternativa para constantes
	var remainingTickets uint = 50

	fmt.Print("Bem vindo a reserva de ingressos da ", conferenceName, ".\n")
	fmt.Printf("Há o total de %v e há %v vagas restantes.\n", referenceTickets, remainingTickets)
	fmt.Println("Obtenha aqui seus passes para participação no", "evento.")

	// := operador de short variable declaration
	// Exemplo: nome := "Andre" ao invés de var nome strng = "Andre"

	// DATA TYPES
	// string
	// int / uint / int8 / int16 ...
	// float32 / float64

	// ARRAYS
	// Tem que definir o tamanho e não pode misturar diferentes tipos
	// var bookings = [50]string{"Andre", "Luiz"}
	// bookings[0] = "Andre"
	// bookings[1] = "Luiz"
	// var bookings [50]string

	// SLICES
	// slice is an abstraction of array
	// slices are more flexible and powerful: variable length or get an sub-array of its own
	// slices are also index-based and have a size, but is resized when needed
	var bookings []string
	// bookings := []string{"Andre", "Luiz"}

	for {
		// Quando não definimos o valor na declaração um tipo dado deve ser informado
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Informe o primeiro nome: ")
		fmt.Scan(&firstName)
		fmt.Print("Informe o último nome: ")
		fmt.Scan(&lastName)
		fmt.Print("Informe o e-mail: ")
		fmt.Scan(&email)
		fmt.Print("Informe a quantidade de ingressos: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Há somente %v ingressos restantes. Você não pode comprar %v\n", referenceTickets, userTickets)
			continue
		}
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf(
			"Obrigado %v %v por reservar %v ingressos. A confirmação será enviada para o e-mail %v.\n",
			firstName,
			lastName,
			userTickets,
			email,
		)
		fmt.Printf("%v ingressos restantes para %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		// for index, booking := range bookings {
		// _ é usado para identificar variáveis não usadas
		for _, booking := range bookings {
			// No moldes do #split no ruby, aceita o separador como segundo parâmetro
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		// "range" iterates over elements for different data structures (so not only array and slices )

		fmt.Printf("Os primeiros nomes dos ingressos: %v\n", firstNames)

		// var noTickerRemaining bool = remainingTickets == 0
		if remainingTickets == 0 {
			fmt.Println("Os ingressos acabaram. Volte novamente ano que vem!")
			break
		}
	}

}
