package main

import (
	"fmt"
	"time"
)

type Title string // book title
type Name string  // library member name

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total owned by the library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBook(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {

	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["The little go book"] = BookEntry{
		total:  3,
		lended: 0,
	}

	library.books["Let's learn go"] = BookEntry{
		total:  2,
		lended: 0,
	}

	library.books["Go bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Mantas"] = Member{"Mantas", make(map[Title]LendAudit)}
	library.members["Rytis"] = Member{"Rytis", make(map[Title]LendAudit)}
	library.members["Lukas"] = Member{"Lukas", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBook(&library)
	printMemberAudits(&library)

	member := library.members["Mantas"]
	checkedOut := checkoutBook(&library, "Go bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBook(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go bootcamp", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBook(&library)
		printMemberAudits(&library)
	}
}
