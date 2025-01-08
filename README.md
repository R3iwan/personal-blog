# Personal Blog Application

This is a simple console-based application for managing a personal blog. It includes two sections: **Guest Section** (for public access) and **Admin Section** (for authenticated users). The application allows for viewing, creating, editing, and deleting blog articles.

---

## Features

### Guest Section
- **Home Page**: View a list of all published articles with their titles and publication dates.
- **Article Page**: View the content of a specific article.

### Admin Section (Requires Authentication)
- **Dashboard**: View a list of articles with options to add, edit, or delete articles.
- **Add Article**: Create and publish a new article by entering the title, content, and date of publication.
- **Edit Article**: Modify the title, content, or date of an existing article.
- **Delete Article**: Remove an article permanently.

---

## Authentication
- **Admin Login**: Authenticate using a hardcoded username (`admin`) and password (`password`).
- **Session Management**: The admin session is tracked, and only authenticated users can access the Admin Section.
- **Logout**: End the admin session at any time.

---

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/r3iwan/personal-blog.git
   cd personal-blog
2. Build and run the application:
   ```bash
   go run main.go
