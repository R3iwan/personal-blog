package guest

import (
	"github.com/r3iwan/personal-blog/pkg/models"
	"github.com/r3iwan/personal-blog/pkg/service"
)

func HomePage(savedArticles *models.SavedArticles) {
	service.ViewHomepage(savedArticles)
}

func ArticlePage(savedArticles *models.SavedArticles) {
	service.ViewArticle(savedArticles)
}
