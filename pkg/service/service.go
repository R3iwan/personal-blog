package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/r3iwan/personal-blog/pkg/models"
)

func ViewHomepage(savedArticles *models.SavedArticles) {
	if EmptyArticles(savedArticles) {
		return
	}

	fmt.Println("--Personal Blog--")
	for _, article := range *savedArticles {
		fmt.Printf("%s \t\t %s\n", article.Title, article.Date)
	}
}

func ViewArticle(savedArticles *models.SavedArticles) {
	if EmptyArticles(savedArticles) {
		return
	}

	var id int
	fmt.Scan(&id)
	found := false

	for _, article := range *savedArticles {
		if article.ID == id {
			fmt.Printf("%s\n%s\n%s\n", article.Title, article.Date, article.Content)
			found = true
			break
		}
	}

	if found {
		fmt.Println("Articles are not found")
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

func EmptyArticles(savedArticles *models.SavedArticles) bool {
	if len(*savedArticles) == 0 {
		fmt.Println("Articles are empty")
		return true
	}
	return false
}

