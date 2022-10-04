package main

import (
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Conferência Go"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
		)

		if isValidName && isValidEmail && isValidTicketsNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("Os primeiros nomes dos ingressos: %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("Bem-vindo a apricação de venda de ingressas da %v!\n", conferenceName)
	fmt.Printf("Há o total de %v e há %v vagas restantes.\n", conferenceTickets, remainingTickets)
	fmt.Println("Obtenha aqui seus passes para participação no", "evento.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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
	userTickets uint,
	firstName string,
	lastName string,
	email string,
) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf(
		"Obrigado %v %v por reservar %v ingressos. A confirmação será enviada para o e-mail %v.\n",
		firstName,
		lastName,
		userTickets,
		email,
	)
	fmt.Printf("%v ingressos restantes para %v\n", remainingTickets, conferenceName)
}
