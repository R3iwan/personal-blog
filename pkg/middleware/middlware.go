package middleware

import (
	"fmt"

	"github.com/r3iwan/personal-blog/pkg/auth"
)

func RequireAdminAuth() {
	if !auth.IsAuthenticated() {
		fmt.Println("Access denied. Please log in as admin.")
		return
	}
}