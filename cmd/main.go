package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/r3iwan/personal-blog/pkg/admin"
	"github.com/r3iwan/personal-blog/pkg/auth"
	"github.com/r3iwan/personal-blog/pkg/guest"
	"github.com/r3iwan/personal-blog/pkg/models"
	"github.com/r3iwan/personal-blog/pkg/service"
)

func main() {
	var savedArticles models.SavedArticles
	service.LoadArticlesFromJSON("articles.json", &savedArticles)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nOptions: 1. Guest Section 2. Admin Section 3. Login 4. Logout 5. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			guest.HomePage(&savedArticles)

		case "2":
			if auth.IsAuthenticated() {
				admin.Dashboard(&savedArticles)
			} else {
				fmt.Println("Access denied. Please log in as admin.")
			}

		case "3":
			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			if err := auth.AuthAdmin(username, password); err != nil {
				fmt.Println("Login failed:", err)
			} else {
				fmt.Println("Login successful!")
			}

		case "4":
			auth.LogoutAdmin()

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
