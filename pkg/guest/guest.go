package guest

import (
	"fmt"

	"github.com/r3iwan/personal-blog/pkg/models"
	"github.com/r3iwan/personal-blog/pkg/service"
)

func GuestSection() {
	var savedArticles models.SavedArticles

	err := service.LoadArticlesFromJSON("articles.json", &savedArticles)
	if err != nil {
		fmt.Println("Error loading articles:", err)
	}

	for {
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			service.AddArticle(&savedArticles)
		case 2:
			service.ViewHomepage(&savedArticles)
		}
	}
}
