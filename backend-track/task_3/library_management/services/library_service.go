package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(title, author string)
	AddMember(name string)
	RemoveBook(bookId int)
	BorrowBook(bookId int, memberId int) error
	ReturnBook(bookId int, memberId int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
	ListMemebers() []models.Member
}

type Library struct {
	books        map[int]models.Book
	members      map[int]models.Member
	nextBookId   int
	nextMemberId int
}

func NewLibrary() *Library {
	newLibrary := Library{
		books:        make(map[int]models.Book),
		members:      make(map[int]models.Member),
		nextBookId:   1,
		nextMemberId: 1,
	}
	return &newLibrary
}

// implemenet the library manager interface for the library struct

func (l *Library) AddBook(title, author string) {
	book := models.Book{
		ID:     l.nextBookId,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	l.books[l.nextBookId] = book
	l.nextBookId++
}
func (l *Library) AddMember(name string) {
	newMember := models.Member{ID: l.nextMemberId, Name: name, BorrowedBooks: []models.Book{}}
	l.members[newMember.ID] = newMember
	l.nextMemberId++
}

func (l *Library) RemoveBook(id int) error {
	book, book_exists := l.books[id]
	if !(book_exists) {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book not in the library")
	}
	delete(l.books, id)
	return nil
}

func (l *Library) BorrowBook(bookId int, memberId int) error {
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

func (l *Library) ReturnBook(bookId int, memeberId int) error {
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

func (l *Library) ListAvailableBooks() []models.Book {

	var found_books []models.Book

	for _, book := range l.books {

		if book.Status == "Available" {
			found_books = append(found_books, book)
		}
	}

	return found_books

}

func (l *Library) ListBorrowedBooks(memberId int) []models.Book {

	member, exists := l.members[memberId]
	if !exists {
		return nil
	}
	return member.BorrowedBooks

}

func (l *Library) ListMemebers() []models.Member {
	var member_list []models.Member
	for _, value := range l.members {
		member_list = append(member_list, value)
	}
	return member_list
}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// /kbjsdbjbjfs/jkdfbjkgkdfbjdf jkdfjgjksdjbfjkfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jhhjsdhjfhjsfkdfjgjksdjbfjksjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdhjbjhsdjhbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdjhjksdjk
//jkdfbjkgkdfbjdf sdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdfdhjjhsdhjjhsdjkdfjgdfhbjhjfdhjfhjdsksdjbfjksdfjsjdhjsdjhjksdjk
// /jsdjhjksdjk
//jkdfbjkgkdfjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdsdjhjfshjhjfshjjksdjk
//dhjjhsdjhjhsdkdfbjkgkdfbjdf jkdfjgjksdhjjhsjhhjshjbfksdbfjsjddhjbjshdhjfshjdjhjsdjhjksdjk
// hjhj
