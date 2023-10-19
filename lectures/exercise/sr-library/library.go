//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

/*
Struct livros
//Struct members
funçao de emprestimo (livro + membro) (horario de emprestimo)
função de entrega (livro + membro) (horario de emprestimo)
função de verificar o status de todos os livros
*/

// Membros
const (
	membroPadrao = 0
	alessandra   = 1
	bianca       = 2
	danny        = 3
)

// Livros
const (
	guia1   = 0
	ecce    = 1
	fredo   = 2
	candido = 3
)

type Book struct {
	index             int
	nome              string
	emprestado        bool
	emprestadoPorQuem Member
	emprestadoQuando  time.Time
}

type Member struct {
	nome        string
	quantLivros int
	livros      []Book
}

func checarLib(books []*Book, clients []*Member) {
	fmt.Println("Status da lib")
	fmt.Println("---")
	for _, book := range books {
		if !book.emprestado {
			fmt.Println("Nome:", book.nome, "/ Emprestado? Não")
		} else {
			fmt.Println("Nome:", book.nome, "/ Emprestado? Sim")
		}
	}

	fmt.Println("---")
	//fmt.Println("Membros com livros emprestados")
	for _, client := range clients {
		if client.quantLivros > 0 {
			//fmt.Println(client.nome, "/ Livros emprestados", client.quantLivros)
			checarClients(client)
		}
	}
	//fmt.Println("---")
}

func checarClients(client *Member) {
	fmt.Println("Livros em posse de", client.nome, ":")
	for _, book := range client.livros {
		fmt.Println(book.nome, ", desde", book.emprestadoQuando.Format("2006-01-02 15:04"))
	}
	fmt.Println("---")
}

func emprestarLivro(book *Book, client *Member) {
	fmt.Println("Emprestando livro...")
	book.emprestado = true
	book.emprestadoPorQuem = *client
	book.emprestadoQuando = time.Now()
	formattedTime := book.emprestadoQuando.Format("2006-01-02 15:04")
	client.quantLivros++
	client.livros = append(client.livros, *book)
	fmt.Println("Livro:", book.nome, "Emprestado para", client.nome, "em:", formattedTime)

	fmt.Println("---")
}

func devolverLivro(book *Book, client *Member, membros []*Member) {
	fmt.Println("Devolvendo livro...")
	fmt.Println("Livro:", book.nome, "Devolvido por", client.nome, "em:", time.Now().Format("2006-01-02 15:04"))

	book.emprestado = false
	client.quantLivros--
	client.quantLivros[book.index]
	client = membros[membroPadrao]
	book.emprestadoPorQuem = *client

	fmt.Println("---")
}

func main() {

	//Livros

	books := []*Book{
		{nome: "O guia do mochileiro das galaxias", index: guia1},
		{nome: "Ecce Homo", index: ecce},
		{nome: "Fredo", index: fredo},
		{nome: "Cândido", index: candido},
	}

	membros := []*Member{
		{nome: "membroPadrao"},
		{nome: "Alessandra Ribeiro"},
		{nome: "Bianca Hills"},
		{nome: "Danny Bendochy"},
	}

	checarLib(books, membros)

	emprestarLivro(books[guia1], membros[alessandra])

	checarLib(books, membros)

	devolverLivro(books[guia1], membros[alessandra], membros)

	checarLib(books, membros)

	emprestarLivro(books[ecce], membros[danny])
	emprestarLivro(books[fredo], membros[danny])
	emprestarLivro(books[candido], membros[bianca])

	checarLib(books, membros)

	devolverLivro(books[ecce], membros[danny], membros)

	checarLib(books, membros)
}
