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

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[Not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
		//fmt.Println("1")
	}
}

func printMembersAudit(library *library) {
	for _, member := range library.members {
		printMemberAudit(&member)
		//fmt.Println("0")
	}
}

func printLibraryBooks(library *library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ Lended", book.lended)
	}
}

func checkOutBook(library *library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book no part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No more books to lend")
		return false
	}

	book.lended++
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}

	fmt.Println("Checked out", title, "to", member.name)
	return true
}

func returnBook(library *library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book no part of library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended--
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func main() {

	library := library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["O guia do Mochileiro das Galaxias"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Ecce Homo"] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.books["Candido"] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.books["Fredo"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Alessandra"] = Member{"Alessandra Ribeiro", make(map[Title]LendAudit)}
	library.members["Danny"] = Member{"Danny Bendochy", make(map[Title]LendAudit)}
	library.members["Juliana Souza"] = Member{"Juliana Souza", make(map[Title]LendAudit)}

	fmt.Println("Initials")
	printLibraryBooks(&library)
	printMembersAudit(&library)

	member := library.members["Alessandra"]
	fmt.Println("Checkout a book")
	checkOutBook := checkOutBook(&library, "Ecce Homo", &member)
	if checkOutBook {
		printLibraryBooks(&library)
		printMembersAudit(&library)
	} else {
		fmt.Println("Error")
	}

}
