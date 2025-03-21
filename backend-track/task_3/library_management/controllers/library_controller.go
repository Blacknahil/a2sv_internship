package controllers

import (
	"bufio"
	"fmt"
	"library_management/services"
	"os"
)

func RunLibrarySystem() {

	library := services.NewLibrary()
	reader := bufio.NewReader(os.Stdin)
	var choice int

	for {
		fmt.Println("\033[34m \n \n Welcome to Library Management console app: \033[0m")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. List Members")
		fmt.Println("7. List Available Books")
		fmt.Println("8. List Borrowed Books")
		fmt.Println("9. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title, author string
			fmt.Print("Enter title: ")
			title, _ = reader.ReadString('\n')
			fmt.Print("Enter author: ")
			author, _ = reader.ReadString('\n')
			library.AddBook(title, author)
			fmt.Println("\033[33m Book added successfully!\033[0m")
		case 2:
			var member_name string
			fmt.Print("Enter name: ")
			member_name, _ = reader.ReadString('\n')
			library.AddMember(member_name)
			fmt.Println("\033[33m Member added successfully!\033[0m")

		case 3:
			var id int
			fmt.Print("Enter book ID to remove: ")
			fmt.Scanln(&id)
			if err := library.RemoveBook(id); err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("\033[33m Book removed successfully!\033[0m")
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scanln(&memberID)
			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(" \033[33m Book borrowed successfully! \033[0m ")
			}
		case 5:

			var bookID, memberID int
			fmt.Print("Enter book ID to return: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scanln(&memberID)
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("\033[31m Error: \033[0m ", err)
			} else {
				fmt.Println("\033[33m Book returned successfully!\033[0m")
			}
		case 6:
			members := library.ListMemebers()
			fmt.Println("\033[33m Library members  \033[0m", members)

		case 7:
			books := library.ListAvailableBooks()
			fmt.Println("\033[33m Available Books:  \033[0m ", books)
		case 8:
			var memberID int
			fmt.Print("Enter member ID: ")
			fmt.Scanln(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("\033[33m Borrowed Books: \033[0m ", books)
		case 9:
			fmt.Println(" \033[35m Existing ...... Goodbye\033[0m")
			return
		}
	}

}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjsjhfshjhjjhjksdjk
//jkdfbjkgkdfbjdf jkdfhjsdhjjhslgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjjdhfhjsdjhhjs f jkdfjghjfjhhjsfjksdjbfjksdbfjsjdhjsdjhjksdjk
//kgkdfbjdf jkdfjgjksdsjdhbjhsdhhjjbjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbjfhjhfdsjhjhfdsksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbdjbhjsdjskjdjkjsjdhjsdjhjksdjk
// /jhjsjhjjkdfbjkgkdfbjdf jjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jsdjhjksdhjk
//jkdfbjkgkdfbjdf jsdjkkj
//dfhjsdhjjhkdfjgjksdjbfjjshjhjjhjkk
//jkdfbjkgkdfbjdf jkdfjksdjk
// /fjkdjbkjkdfj/jhlshdjjhsdfhjhjdsfjksdjk
// jmbxcvhjdjh
// \033[33m \033[0m
// hjdsfjhdhjsjh
