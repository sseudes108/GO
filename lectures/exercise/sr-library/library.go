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

const (
	membroPadrao = 0
	alessandra   = 1
	bianca       = 2
	danny        = 3
)

type Book struct {
	nome              string
	emprestado        bool
	emprestadoPorQuem Member
	retiradoQuando    time.Time
}

type Member struct {
	nome              string
	livrosEmprestados int
}

func statusDosLivrosNaLib(books []*Book, clients []*Member) {
	fmt.Println("Status da lib")
	fmt.Println("---")
	for _, book := range books {
		if !book.emprestado {
			fmt.Println("Nome:", book.nome, "/ Emprestado? Sim")
		} else {
			fmt.Println("Nome:", book.nome, "/ Emprestado? Não")
		}
	}

	fmt.Println("---")
	fmt.Println("Membros com livros emprestados")
	for _, client := range clients {
		if client.livrosEmprestados > 0 {
			fmt.Println(client.nome, "/ Livros emprestados", client.livrosEmprestados)
		}
	}

	fmt.Println("---")
}

func emprestarLivro(book *Book, client *Member) {
	book.emprestado = true
	book.emprestadoPorQuem = *client
	book.retiradoQuando = time.Now()
	formattedTime := book.retiradoQuando.Format("2006-01-02 15:04")
	client.livrosEmprestados++
	fmt.Println("Livro:", book.nome, "Emprestado para", client.nome, "em:", formattedTime)

	fmt.Println("---")
}

func devolverLivro(book *Book, client *Member, membros []*Member) {

	fmt.Println("Livro:", book.nome, "Devolvido por", client.nome, "em:", time.Now().Format("2006-01-02 15:04"))

	book.emprestado = false
	client = membros[membroPadrao]
	book.emprestadoPorQuem = *client

	fmt.Println("---")
}

func main() {

	//Livros
	books := []*Book{
		{nome: "O guia do mochileiro das galaxias"},
		{nome: "Ecce Homo"},
		{nome: "Fredo"},
		{nome: "Cândido"},
	}

	membros := []*Member{
		{nome: "membroPadrao"},
		{nome: "Alessandra Ribeiro", livrosEmprestados: 0},
		{nome: "Bianca Hills", livrosEmprestados: 0},
		{nome: "Danny Bendochy", livrosEmprestados: 0},
	}

	statusDosLivrosNaLib(books, membros)

	emprestarLivro(books[0], membros[alessandra])

	statusDosLivrosNaLib(books, membros)

	devolverLivro(books[0], membros[alessandra], membros)

	emprestarLivro(books[2], membros[danny])
	emprestarLivro(books[3], membros[bianca])

	statusDosLivrosNaLib(books, membros)
}
