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

// Registro dos horarios de retirada e devolução
type Emprestimo struct {
	retirada  time.Time
	devolucao time.Time
}

type Membro struct {
	nome   Nome
	livros map[Titulo]Emprestimo
}

// Registro de quantidades do Livro em questão
type LivroReg struct {
	total       int
	emprestados int
}

type Biblioteca struct {
	membros map[Nome]Membro
	livros  map[Titulo]LivroReg
}

func verificarEmprestimos(membro *Membro) {
	for titulo, emprestimo := range membro.livros {
		var retornoTempo string
		if emprestimo.devolucao.IsZero() {
			retornoTempo = "[Ainda não devolvido]"
		} else {
			retornoTempo = emprestimo.devolucao.String()
		}
		fmt.Println(membro.nome, ":", titulo, ":", emprestimo.retirada.String(), "até????", retornoTempo)
	}
}

func checarMembros(bib *Biblioteca) {
	for _, membro := range bib.membros {
		verificarEmprestimos(&membro)
	}
}

func checarBiblioteca(bib *Biblioteca) {
	fmt.Println()
	for titulo, livro := range bib.livros {
		fmt.Println(titulo, "/ Total:", livro.total, "/ Emprestados:", livro.emprestados,
			"/ Disponiveis:", livro.total-livro.emprestados)
	}
}

func emprestarLivro(bib *Biblioteca, titulo Titulo, membro *Membro) bool {
	fmt.Println("\n", titulo)
	livro, encontrado := bib.livros[titulo]
	if !encontrado {
		fmt.Println("Livro não faz parte da biblioteca")
		return false
	}
	if livro.emprestados >= livro.total {
		livro.emprestados = livro.total
		fmt.Println("Não há mais desse livro disponivel na biblioteca")
		return false
	}
	livro.emprestados++
	bib.livros[titulo] = livro
	membro.livros[titulo] = Emprestimo{retirada: time.Now()}

	fmt.Println("Livro retirado com sucesso, por", membro.nome)
	return true
}

func devolverLivro(bib *Biblioteca, titulo Titulo, membro *Membro) bool {
	fmt.Println("\n", titulo)
	livro, encontrado := bib.livros[titulo]
	if !encontrado {
		fmt.Println("Livro não faz parte da biblioteca")
		return false
	}

	emprestimo, encontrado := membro.livros[titulo]
	if !encontrado {
		fmt.Println(membro.nome, "não retirou este livro")
		return false
	}

	livro.emprestados--
	bib.livros[titulo] = livro

	emprestimo.devolucao = time.Now()
	membro.livros[titulo] = emprestimo

	fmt.Println("Livro devolvido por", membro.nome)
	return true
}

func main() {

	////// 	BIBILIOTECA //////

	bib := Biblioteca{
		livros:  make(map[Titulo]LivroReg),
		membros: make(map[Nome]Membro),
	}
	////// LIVROS //////
	bib.livros["Ecce Homo"] = LivroReg{
		total:       2,
		emprestados: 0,
	}
	bib.livros["Zaratustra"] = LivroReg{
		total:       2,
		emprestados: 0,
	}
	bib.livros["Gaia Ciencia"] = LivroReg{
		total:       1,
		emprestados: 0,
	}
	bib.livros["Anticristo"] = LivroReg{
		total:       3,
		emprestados: 0,
	}
	////// MEMBROS //////
	bib.membros["Eudes"] = Membro{"Eudes", make(map[Titulo]Emprestimo)}
	bib.membros["Alessandra"] = Membro{"Alessandra", make(map[Titulo]Emprestimo)}
	bib.membros["Bruna"] = Membro{"Bruna", make(map[Titulo]Emprestimo)}

	fmt.Println("INICIAIS")
	checarBiblioteca(&bib)
	checarMembros(&bib)

	alessandra := bib.membros["Alessandra"]
	emprestarLivro(&bib, "Zaratustra", &alessandra)
	emprestarLivro(&bib, "Anticristo", &alessandra)

	bruna := bib.membros["Bruna"]
	emprestarLivro(&bib, "Zaratustra", &bruna)
	emprestarLivro(&bib, "Zaratustra", &bruna)

	devolverLivro(&bib, "Zaratustra", &bruna)
	devolverLivro(&bib, "Zaratustra", &alessandra)
	devolverLivro(&bib, "Anticristo", &alessandra)

	fmt.Println("FINAIS")
	checarMembros(&bib)
	checarBiblioteca(&bib)
}
