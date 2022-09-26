package main

import "fmt"

func main() {
	// Go infere o tipo de dado quando definimos o valor na declaração
	var conferenceName = "Conferência Go" // alternativa => conferenceName := 50, e não funciona pra declarar o tipo explicitamente
	const referenceTickets = 50           // não há alternativa para constantes
	var remainingTickets uint = 50

	fmt.Print("Bem vindo a reserva de ingressos da ", conferenceName, ".\n")
	fmt.Printf("Há o total de %v e há %v vagas restantes.\n", referenceTickets, remainingTickets)
	fmt.Println("Obtenha aqui seus passes para participação no", "evento.")

	// Data types
	// string
	// int / uint / int8 / int16 ...
	// float32 / float64

	// Tem que definir o tamanho e não pode misturar diferentes tipos
	// var bookings = [50]string{"Andre", "Luiz"}
	// bookings[0] = "Andre"
	// bookings[1] = "Luiz"

	var bookings [50]string

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

	bookings[0] = firstName + lastName

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Array inteiro: %v\n", bookings)
	fmt.Printf("Array primeiro item: %v\n", bookings[0])
	fmt.Printf("Array tipo: %T\n", bookings[0])
	fmt.Printf("Array tamanho: %v\n", len(bookings))

	fmt.Printf(
		"Obrigado %v %v por reservar %v ingressos. A confirmação será enviada para o e-mail %v.\n",
		firstName,
		lastName,
		userTickets,
		email,
	)

	fmt.Printf("%v ingressos restantes para %v\n", remainingTickets, conferenceName)
}
