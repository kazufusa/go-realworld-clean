package domain

import (
	"time"
)

type Article struct {
	Slug        string
	Title       string
	Description string
	Body        string
	TagList     []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FavoritedBy []User
	Author      User
	Comments    []Comment
}

type Comment struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	Author    User
}
type ArticleUpdatableField int

const (
	Title ArticleUpdatableField = iota
	Description
	Body
)

func UpdateArticle(initial *Article, opts ...func(fields *Article)) {
	for _, v := range opts {
		v(initial)
	}
}

func SetArticleTitle(input *string) func(fields *Article) {
	return func(initial *Article) {
		if input == nil {
			return
		}
		initial.Title = *input
	}
}

func SetArticleDescription(input *string) func(fields *Article) {
	return func(initial *Article) {
		if input == nil {
			return
		}
		initial.Description = *input
	}
}

func SetArticleBody(input *string) func(fields *Article) {
	return func(initial *Article) {
		if input == nil {
			return
		}
		initial.Body = *input
	}
}

type ArticleFilter func(Article) bool

func ArticleHasTag(tag string) ArticleFilter {
	return func(article Article) bool {
		for _, articleTag := range article.TagList {
			if articleTag == tag {
				return true
			}
		}
		return false
	}
}

func ArticleHasAuthor(authName string) ArticleFilter {
	return func(article Article) bool {
		return article.Author.Name == authName
	}
}

func ArticleIsFavoritedBy(username string) ArticleFilter {
	return func(article Article) bool {
		if username == "" {
			return false
		}
		for _, user := range article.FavoritedBy {
			if user.Name == username {
				return true
			}
		}
		return false
	}
}

type ArticleCollection []Article

func (articles ArticleCollection) ApplyLimitAndOffset(limit, offset int) ArticleCollection {
	if limit <= 0 {
		return []Article{}
	}

	articleSize := len(articles)
	min := offset
	if min < 0 {
		min = 0
	}
	if articleSize < min {
		return []Article{}
	}

	max := min + limit
	if articleSize < max {
		max = articleSize
	}

	return articles[min:max]
}

func (article *Article) UpdateComments(comment Comment, add bool) {
	if add {
		article.Comments = append(article.Comments, comment)
		return
	} else {
		for i, c := range article.Comments {
			if c.ID == comment.ID {
				article.Comments = append(article.Comments[0:i], article.Comments[i+1:len(article.Comments)]...)
			}
		}
	}
}

func (article *Article) UpdateFavoritedBy(user User, add bool) {
	if add {
		for _, favocitedBy := range article.FavoritedBy {
			if favocitedBy.Name == user.Name {
				return
			}
		}
		article.FavoritedBy = append(article.FavoritedBy, user)
		return
	} else {
		for i, favocitedBy := range article.FavoritedBy {
			if favocitedBy.Name == user.Name {
				article.FavoritedBy = append(article.FavoritedBy[0:i], article.FavoritedBy[i+1:len(article.FavoritedBy)]...)
			}
		}
	}
}
