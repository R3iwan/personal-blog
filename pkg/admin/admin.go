package admin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/r3iwan/personal-blog/pkg/middleware"
	"github.com/r3iwan/personal-blog/pkg/models"
	"github.com/r3iwan/personal-blog/pkg/service"
)

var mu sync.Mutex

func Dashboard(savedArticles *models.SavedArticles) {
	middleware.RequireAdminAuth()
	service.ViewHomepage(savedArticles)
	fmt.Println("Options: 1. Add Article 2. Edit Article 3. Delete Article")
}

func AddArticlePage(savedArticles *models.SavedArticles) {
	middleware.RequireAdminAuth()
	AddArticle(savedArticles)
}

func EditArticlePage(savedArticles *models.SavedArticles) {
	middleware.RequireAdminAuth()
	EditArticle(savedArticles)
}

func DeleteArticlePage(savedArticles *models.SavedArticles) {
	middleware.RequireAdminAuth()
	DeleteArticle(savedArticles)
}

func AddArticle(savedArticles *models.SavedArticles) {
	title, content, date := WriteInfo()

	nextID := 1
	if len(*savedArticles) > 0 {
		nextID = (*savedArticles)[len(*savedArticles)-1].ID + 1
	}

	mu.Lock()
	article := models.Article{
		ID:      nextID,
		Title:   title,
		Date:    date,
		Content: content,
	}

	*savedArticles = append(*savedArticles, article)

	err := service.SaveArticleToJSON("articles.json", models.SavedArticles{article})
	mu.Unlock()

	if err != nil {
		fmt.Println("Error saving articles to JSON file...")
	}

}

func EditArticle(savedArticles *models.SavedArticles) {
	if service.EmptyArticles(savedArticles) {
		return
	}

	var id int
	fmt.Println("Enter the ID of the article to edit:")
	fmt.Scan(&id)
	found := false

	for i, article := range *savedArticles {
		if article.ID == id {
			title, content, date := WriteInfo()
			(*savedArticles)[i].Title = title
			(*savedArticles)[i].Date = date
			(*savedArticles)[i].Content = content

			err := service.SaveArticleToJSON("articles.json", *savedArticles)
			if err != nil {
				fmt.Println("Error editing article")
			}

			fmt.Println("Article updated successfully!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Articles with ID not found")
	}
}

func DeleteArticle(savedArticles *models.SavedArticles) {
	if service.EmptyArticles(savedArticles) {
		return
	}

	var id int
	fmt.Println("Enter the ID of the article to edit:")
	fmt.Scan(&id)
	found := false

	for i, article := range *savedArticles {
		if article.ID == id {
			*savedArticles = append((*savedArticles)[:i], (*savedArticles)[i+1:]...)

			err := service.SaveArticleToJSON("articles.json", *savedArticles)
			if err != nil {
				fmt.Println("Error saving updated articles to JSON file:", err)
				return
			}

			fmt.Println("Article updated successfully!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Articles with ID not found")
	}
}

func WriteInfo() (string, string, string) {
	date := time.Now().Format("January 2, 2006. Time: 15:04:05")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the title of the article:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Please enter the content of the article:")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	return title, content, date
}
