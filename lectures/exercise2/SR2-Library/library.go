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

type Nome string

type Time struct {
	retirado  time.Time
	devolvido time.Time
}

type Livro struct {
	nome        Nome
	quantidade  int
	alugados    int
	retiradoEm  Time
	devolvidoEm Time
}

type Membro struct {
	nome     Nome
	alugados int
}

/*type Biblioteca struct {
	livros  []*Livro
	members []*Membro
}*/

type Biblioteca struct {
	livros  map[Nome]*Livro
	members []*Membro
}

func criarBiblioteca() *Biblioteca {

	lib := Biblioteca{
		livros:=make(map[Nome]Livro)
	}
	return lib
}

func (lib *Biblioteca) emprestarLivro(livro *Livro, client *Membro) {
	fmt.Println(lib.livros[livro].nome)
}

func (lib *Biblioteca) infoLivros() {
	fmt.Println()
	fmt.Println("Livros")
	for _, livro := range lib.livros {
		fmt.Println(livro.nome, " - Quantidade em estante:", livro.quantidade,
			" - Quantidade emprestada:", livro.alugados)
	}
}

func (lib *Biblioteca) infoMembros() {
	fmt.Println()
	fmt.Println("Membros")
	for _, membro := range lib.members {
		fmt.Println(membro.nome, "- Alugados:", membro.alugados)
	}
}

func (lib *Biblioteca) infoLib() {
	lib.infoLivros()
	lib.infoMembros()
}

func main() {
	lib := criarBiblioteca()

	//fmt.Println(lib)

	lib.infoLivros()
	lib.infoMembros()
}
