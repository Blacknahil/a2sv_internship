# Library Management System

Welcome to the **Library Management System**, a console-based application for managing books and members in a library. This system allows users to perform various operations such as adding books, managing members, borrowing and returning books, reserving books, and viewing available or borrowed books.

---

## Features

The Library Management System provides the following functionalities:

1. **Add Book**: Add a new book to the library by providing its title and author.
2. **Add Member**: Register a new member in the library by providing their name.
3. **Remove Book**: Remove a book from the library (only if it is not currently borrowed or reserved).
4. **Return Book**: Return a borrowed book to the library by providing the book ID and member ID.
5. **Borrow Book**: Borrow a book for a specific member by providing the book ID and member ID.
6. **Reserve Book**: Reserve a book for a specific member. If the book is not borrowed within 5 seconds, the reservation is automatically canceled.
7. **List Members**: View a list of all registered members in the library.
8. **List Available Books**: View a list of all books currently available in the library.
9. **List Borrowed Books**: View a list of all books borrowed by a specific member.
10. **Exit**: Exit the application.

---

## Concurrency Features

- **Concurrent Book Reservation**: Multiple members can attempt to reserve books simultaneously, managed safely using Goroutines and Channels.
- **Auto-Cancellation of Reservations**: If a book is reserved but not borrowed within 5 seconds, it becomes available again.
- **Thread-Safe Operations**: The system ensures safe concurrent access using `sync.Mutex` to avoid race conditions.

---

## How to Use

1. Clone this repository to your local machine.
2. Navigate to the `library_management` directory.
3. Run the application using the following command:
   ```bash
   go run main.go
   ```