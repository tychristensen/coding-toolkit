package book_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/fxtlabs/date"
	"github.com/tydanny/coding-toolkit/book"
)

// This is the main test area
var _ = Describe("Book", func() {
	var lotr, harPot, enderGame *book.Book
	var books book.ByPubDate

	// Setup books to test with
	BeforeEach(func() {
		lotr = &book.Book{
			Title:   "Lord of the Rings: The Fellowship of the Ring",
			Author:  "J. R. R. Tolkien",
			PubDate: date.New(1954, time.July, 29),
		}

		harPot = &book.Book{
			Title:   "Harry Potter and the Sorcerer's Stone",
			Author:  "J. K. Rowling",
			PubDate: date.New(1997, time.June, 26),
		}

		enderGame = &book.Book{
			Title:   "Ender's Game",
			Author:  "Orson Scott Card",
			PubDate: date.New(1985, time.January, 15),
		}

		books = book.ByPubDate{*enderGame, *harPot, *lotr}
	})

	Describe("Sorting Books", func() {
		Context("in ascending order", func() {
			It("should be sorted", func() {
				books.SortAscending()
				Expect(isSorted(books, ascending)).To(BeTrue())
			})
		})

		Context("in descending order", func() {
			It("should be sorted", func() {
				books.SortDesceding()
				Expect(isSorted(books, descending)).To(BeTrue())
			})
		})
	})
})

// sortMode is used to specify what mode to sort in
type sortMode int

const (
	ascending  sortMode = 0
	descending sortMode = 1
)

// isSorted checks to see if books is sorted
// using the proveded compare function
func isSorted(books []book.Book, mode sortMode) bool {
	// Check to see if all books have been
	// sorted in ascending order
	for i, book := range books {
		// Don't check the last book in the slice
		if i == len(books)-1 {
			break
		}

		// TODO: There is probably a better way to do this
		switch mode {
		case ascending:
			if book.PubDate.After(books[i+1].PubDate) {
				return false
			}
		case descending:
			if book.PubDate.Before(books[i+1].PubDate) {
				return false
			}
		}
	}
	return true
}
