package main

import (
	"fmt"
	"strings"
)

func main() {
	// DECLARAÇÃO DE VARIÁVEIS
	// Go infere o tipo de dado quando definimos o valor na declaração
	var conferenceName = "Conferência Go" // alternativa => conferenceName := 50, e não funciona pra declarar o tipo explicitamente
	const conferenceTickets = 50          // não há alternativa para constantes
	var remainingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

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

	// for remainingTickets > 0 {
	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketsNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("Os primeiros nomes dos ingressos: %v\n", firstNames)

			// var noTickerRemaining bool = remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("Os ingressos acabaram. Volte novamente ano que vem!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("O nome informado é muito curto.")
			}
			if !isValidEmail {
				fmt.Println("O e-mail informado é inválido.")
			}
			if !isValidTicketsNumber {
				fmt.Println("A quantidade de e-mails informada não é válida.")
			}
		}
	}
}

func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Bem-vindo a apricação de venda de ingressas da %v!\n", confName)
	fmt.Printf("Há o total de %v e há %v vagas restantes.\n", conferenceTickets, remainingTickets)
	fmt.Println("Obtenha aqui seus passes para participação no", "evento.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
	remainingTickets uint,
) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketsNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(
	remainingTickets uint,
	userTickets uint,
	bookings []string,
	firstName string,
	lastName string,
	email string,
	conferenceName string,
) {
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
}
