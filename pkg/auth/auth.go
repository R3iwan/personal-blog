package auth

import "errors"

var adminUsername = "admin"
var adminPassword = "password"
var isAuthenticated = false

func AuthAdmin(username, password string) error {
	if username == adminUsername && password == adminPassword {
		isAuthenticated = true
		return nil
	}
	return errors.New("invalid username or password")
}

func IsAuthenticated() bool {
	return isAuthenticated
}

func LogoutAdmin() {
	isAuthenticated = false
}
