package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	addBook(book models.Book)
	removeBook(bookId int)
	borrowBook(bookId int, memberId int) error
	returnBook(bookId int, memberId int) error
	listAvailableBooks() []models.Book
	listBorrowedBooks(memberId int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	newLibrary := Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
	return &newLibrary
}

// implemenet the library manager interface for the library struct

func (l *Library) addBook(book models.Book) {
	l.books[book.ID] = book
}

func (l *Library) removeBook(id int) {
	delete(l.books, id)
}

func (l *Library) borrowBook(bookId int, memberId int) error {
	// the book may not exist in the library
	book, book_exists := l.books[bookId]
	if !book_exists {
		return errors.New("book not found")
	}
	// the book may be already borrowed
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}
	// the member may not exist
	member, member_exists := l.members[memberId]
	if !member_exists {
		return errors.New("member not found")
	}

	// borrow the book
	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	// change(update) the status of book and memner inside the library
	l.books[book.ID] = book
	l.members[member.ID] = member
	return nil
}

func (l *Library) returnBook(bookId int, memeberId int) error {
	// the book may not exist
	book, book_exists := l.books[bookId]
	if !book_exists {
		return errors.New("book not found")
	}
	// the book is already in the library
	if book.Status == "Available" {
		return errors.New("book is already in the library")
	}

	member, memeber_exists := l.members[memeberId]
	// member does not exist
	if !memeber_exists {
		return errors.New("member not found")
	}

	for i, taken_book := range member.BorrowedBooks {

		//return the boook
		if taken_book.ID == bookId {
			taken_book.Status = "Available"
			l.books[bookId] = taken_book
			// remove from the memeber.BorrowedBooks list
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			l.members[memeberId] = member
			return nil
		}
	}

	return errors.New("book not borrowed by the member")

}

func (l *Library) listAvailableBooks() []models.Book {

	var found_books []models.Book

	for _, book := range l.books {

		if book.Status == "Available" {
			found_books = append(found_books, book)
		}
	}

	return found_books

}

func (l *Library) listBorrowedBooks() []models.Book {

	var borrowed_books []models.Book

	for _, book := range l.books {

		if book.Status == "Borrowed" {
			borrowed_books = append(borrowed_books, book)
		}
	}

	return borrowed_books
}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// /kbjsdbjbjfs/jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jhhjsdhjfhjsfkdfjgjksdjbfjksfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdhjbjhsdjhbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdsdjhjfshjhjfshjjksdjk
//dhjjhsdjhjhsdkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjddhjbjshdhjfshjdjhjsdjhjksdjk
