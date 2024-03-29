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

type Titulo string
type Nome string

type RegistroTempo struct {
	retirada  time.Time
	devolucao time.Time
}

type RegistroLivro struct {
	quantidade int
	alugados   int
}
type Membro struct {
	nome   Nome
	livros map[Titulo]RegistroTempo
}

type Biblioteca struct {
	livros  map[Titulo]RegistroLivro
	membros map[Nome]Membro
}

func (lib *Biblioteca) devolverLivro(titulo Titulo, membro *Membro) bool {
	fmt.Println("Devolvendo livro...")
	livro, encontrado := lib.livros[titulo]
	if !encontrado {
		fmt.Println(titulo, " - Livro não faz parte da biblioteca")
		return false
	}

	emprestimo, encontrado := membro.livros[titulo]
	if !encontrado {
		fmt.Println(membro.nome, "não retirou este livro")
		return false
	}
	livro.alugados--
	lib.livros[titulo] = livro

	emprestimo.devolucao = time.Now()
	membro.livros[titulo] = emprestimo

	momentoEntrega := emprestimo.devolucao
	horaFormatada := momentoEntrega.Format("2006-01-02 15:04")
	horaFormatadaComoTime, _ := time.Parse("2006-01-02 15:04", horaFormatada)

	delete(membro.livros, titulo)
	fmt.Println(titulo, "devolvido com sucesso por", membro.nome, "em", horaFormatadaComoTime)
	return true
}

func (lib *Biblioteca) emprestarLivro(titulo Titulo, membro *Membro) bool {
	fmt.Println("Emprestando livro...")
	livro, encontrado := lib.livros[titulo]
	if !encontrado {
		fmt.Println(titulo, " - Livro não faz parte da biblioteca")
		return false
	}
	if livro.alugados >= livro.quantidade {
		fmt.Println(titulo, " - Não há mais desse livro disponivel")
		return false
	}
	livro.alugados++
	lib.livros[titulo] = livro

	momentoRetirada := time.Now()
	horaFormatada := momentoRetirada.Format("2006-01-02 15:04")
	horaFormatadaComoTime, _ := time.Parse("2006-01-02 15:04", horaFormatada)
	membro.livros[titulo] = RegistroTempo{retirada: horaFormatadaComoTime}

	fmt.Println(titulo, "emprestado com sucesso para", membro.nome, "em", horaFormatadaComoTime)
	return true
}

func (lib *Biblioteca) infoBiblioteca() {
	fmt.Println("Na Estante:")
	lib.infoLivros()
	fmt.Println("Membros com livros emprestados:")
	lib.infoMembros()
}

func (lib *Biblioteca) verificarEmprestimos(membro *Membro) {
	for titulo, livro := range membro.livros {
		fmt.Println("Membro:", membro.nome, ",", titulo, ",", "desde:", livro.retirada)
	}
}

func (lib *Biblioteca) infoLivros() {
	for titulo, livro := range lib.livros {
		fmt.Println(titulo, livro)
	}
}
func (lib *Biblioteca) infoMembros() {
	for _, membro := range lib.membros {
		lib.verificarEmprestimos(&membro)
	}
}

func criarBiblioteca() *Biblioteca {

	lib := Biblioteca{
		livros:  make(map[Titulo]RegistroLivro),
		membros: make(map[Nome]Membro),
	}

	lib.livros["O Guia do mochileiro das Galaxias Vol.1"] = RegistroLivro{quantidade: 3, alugados: 0}
	lib.livros["Assim falava Zaratustra"] = RegistroLivro{quantidade: 1, alugados: 0}
	lib.livros["A República"] = RegistroLivro{quantidade: 2, alugados: 0}
	lib.livros["Ecce Homo"] = RegistroLivro{quantidade: 1, alugados: 0}

	lib.membros["Alessandra"] = Membro{nome: "Alessandra", livros: make(map[Titulo]RegistroTempo)}
	lib.membros["Danny"] = Membro{nome: "Danny", livros: make(map[Titulo]RegistroTempo)}
	lib.membros["Bianca"] = Membro{nome: "Bianca", livros: make(map[Titulo]RegistroTempo)}
	lib.membros["Sheila"] = Membro{nome: "Sheila", livros: make(map[Titulo]RegistroTempo)}

	return &lib
}

func main() {
	lib := criarBiblioteca()
	lib.infoBiblioteca()

	alessandra := lib.membros["Alessandra"]
	//danny := lib.membros["Danny"]
	//sheila := lib.membros["Sheila"]
	//bianca := lib.membros["Bianca"]

	fmt.Println("INICIO")
	lib.emprestarLivro("Ecce Homo", &alessandra)

	lib.devolverLivro("Ecce Homo", &alessandra)

	lib.infoBiblioteca()
	fmt.Println("FIM")
}
