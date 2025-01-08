package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/r3iwan/personal-blog/pkg/models"
)

var nextID = 1
var mu sync.Mutex

func AddArticle(savedArticles *models.SavedArticles) {
	var title, content string
	date := time.Now().Format("January 2, 2006")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the title of the article:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Please enter the content of the article:")
	content, _ = reader.ReadString('\n')
	content = strings.TrimSpace(content)

	article := models.Article{
		ID:      nextID,
		Title:   title,
		Date:    date,
		Content: content,
	}

	*savedArticles = append(*savedArticles, article)

	mu.Lock()
	err := SaveArticleToJSON("articles.json", models.SavedArticles{article})
	mu.Unlock()
	if err != nil {
		fmt.Println("Error saving articles to JSON file...")
	}

	nextID++

}

func ViewHomepage(savedArticles *models.SavedArticles) {
	if len(*savedArticles) == 0 {
		fmt.Println("Articles are empty")
	}

	fmt.Println("--Personal Blog--")
	for _, article := range *savedArticles {
		fmt.Printf("%s \t\t %s\n", article.Title, article.Date)
	}
}

func SaveArticleToJSON(filename string, newArticles models.SavedArticles) error {
	var existingArticles models.SavedArticles

	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&existingArticles); err != nil {
			return fmt.Errorf("error decoding JSON: %w", err)
		}
	}

	existingArticles = append(existingArticles, newArticles...)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(existingArticles); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}

func LoadArticlesFromJSON(filename string, savedArticles *models.SavedArticles) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		*savedArticles = models.SavedArticles{}
		return nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(savedArticles); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}
