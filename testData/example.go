package testData

import (
	"time"

	"github.com/kazufusa/go-realworld-clean/domain"
)

const TokenPrefix = "Token "

var (
	rickBio = "Rick biography string"
	janeImg = "jane img link"

	rick = domain.User{
		Name:      "rick",
		Email:     "rick@example.com",
		Bio:       &rickBio,
		ImageLink: nil,
		Password:  "password",
	}

	jane = domain.User{
		Name:      "johnjacob",
		Email:     "joe@example.com",
		Bio:       nil,
		ImageLink: &janeImg,
		Password:  "password",
	}

	janeArticle = domain.Article{
		Slug:        "how-to-train-your-dragon",
		Title:       "articleTitle",
		Description: "description",
		Body:        "body",
		TagList:     []string{"tagList"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		FavoritedBy: []domain.User{rick},
		Author:      jane,
		Comments: []domain.Comment{
			{ID: 123,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Body:      "commentBody",
				Author:    rick,
			},
		},
	}
)

func User(name string) domain.User {
	switch name {
	case "rick":
		return rick
	default:
		return jane
	}
}

func Article(name string) domain.Article {
	switch name {
	default:
		return janeArticle
	}
}
