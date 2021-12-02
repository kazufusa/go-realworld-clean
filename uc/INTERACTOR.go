package uc

import "github.com/kazufusa/go-realworld-clean/domain"

type interactor struct {
	logger           Logger
	authHandler      AuthHandler
	userRW           UserRW
	articleRW        ArticleRW
	commentRW        CommentRW
	tagRW            TagsRW
	slugger          Slugger
	userValidator    UserValidator
	articleValidator ArticleValidator
}

type Logger interface {
	Log(...interface{})
}

type AuthHandler interface {
	Create(username, email, password string) (token string, err error)
	GetUserName(token string) (username string, err error)
}

type UserRW interface {
	Create(username, email, passwowrd string) (*domain.User, error)
	GetByName(username string) (*domain.User, error)
	GetByEmailAndPassword(email, password string) (*domain.User, error)
	Save(user domain.User) error
}

type ArticleRW interface {
	Create(domain.Article) (*domain.Article, error)
	Save(domain.Article) (*domain.Article, error)
	GetBySlug(slug string) (*domain.Article, error)
	GetByAuthorsNameOrderedByMostRecentAsc(usernames []string) ([]domain.Article, error)
	GetRecentFiltered(filters []domain.ArticleFilter) ([]domain.Article, error)
	Delete(slug string) error
}

type CommentRW interface {
	Create(comment domain.Comment) (*domain.Comment, error)
	GetByID(id int) (*domain.Comment, error)
	Delete(id int) error
}

type TagsRW interface {
	GetAll() ([]string, error)
	Add(newTags []string) error
}

type Slugger interface {
	NewSlug(string) string
}

type UserValidator interface {
	CheckUser(user domain.User) error
}

type ArticleValidator interface {
	BeforeCreationCheck(article *domain.Article) error
	BeforeUpdateCheck(article *domain.Article) error
}
