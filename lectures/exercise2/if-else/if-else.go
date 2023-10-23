//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

type Day int
type Member string

const (
	monday Day = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	admin      Member = "admin"
	manager    Member = "manager"
	contractor Member = "contractor"
	member     Member = "member"
	guest      Member = "guest"
)

func checkAcess(user Member, day Day) {
	if user == admin || user == manager {
		fmt.Println("Acess Granted", user)
	} else if user == contractor && day > friday {
		fmt.Println("Acess Granted", user)
	} else if user == member && day <= friday {
		fmt.Println("Acess Granted", user)
	} else if user == guest && day == monday || day == wednesday || day == friday {
		fmt.Println("Acess Granted", user)
	} else {
		fmt.Println("Acess Denied", user)
	}
}

func main() {
	checkAcess(admin, friday)
	checkAcess(manager, friday)

	checkAcess(contractor, sunday) //denied

	checkAcess(member, saturday)  //denied
	checkAcess(member, wednesday) //granted

	checkAcess(guest, thursday) //denied
}
