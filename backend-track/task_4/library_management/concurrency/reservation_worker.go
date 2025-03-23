package concurrency

import (
	"errors"
	"library_management/models"
	"sync"
	"time"
)

type ReservationWorker struct {
	mu           sync.Mutex
	reservations map[int]chan struct{}
}

func NewReservationWorker() *ReservationWorker {
	return &ReservationWorker{
		reservations: make(map[int]chan struct{}),
	}
}

func (r *ReservationWorker) ReserveBook(bookID int, memberID int, books map[int]models.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	book, exists := books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != "Available" {
		return errors.New("book already reserved or borrowed")
	}

	book.Status = "Reserved"
	books[bookID] = book
	cancelChan := make(chan struct{})
	r.reservations[bookID] = cancelChan

	go func() {
		select {
		case <-time.After(15 * time.Second):
			r.mu.Lock()
			if books[bookID].Status == "Reserved" {
				books[bookID] = models.Book{ID: bookID, Title: book.Title, Author: book.Author, Status: "Available"}
			}
			r.mu.Unlock()
		case <-cancelChan:
		}
	}()

	return nil
}

func (r *ReservationWorker) CancelReservation(bookID int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cancelChan, reserved := r.reservations[bookID]; reserved {
		close(cancelChan)
		delete(r.reservations, bookID)
	}
}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// hjjhhjfd
// dhjjhsdhj
