package uc

import (
	"github.com/kazufusa/go-realworld-clean/domain"
)

type Handler interface {
	ProfileLogic
	UserLogic
	ArticlesLogic
	ArticleLogic
	CommentsLogic
	FavoritesLogic
	TagsLogic
}

type ProfileLogic interface {
	ProfileGet(requestingUserName, userName string) (profile *domain.User, follows bool, err error)
	ProfileUpdateFollow(requestingUserName, userName string, follow bool) (user *domain.User, err error)
}

type UserLogic interface {
	UserCreate(username, email, password string) (user *domain.User, token, string, err error)
	UserLogin(email, password string) (user *domain.User, token string, err error)
	UserGet(userName string) (user *domain.User, token string, err error)
	UserEdit(userName string, fieldsToUpdate map[domain.UserUpdatableProperty]*string) (user *domain.User, token string, err error)
}

type ArticlesLogic interface {
	ArticlesFeed(username string, limit, offset int) (requestingUser *domain.User, articles domain.ArticleCollection, totalArticleCount int, err error)
	GetArticles(username string, limit, offset int, filter []domain.ArticleFilter) (requestingUser *domain.User, articles domain.ArticleCollection, totalArticleCount int, err error)
}

type ArticleLogic interface {
	ArticleGet(slug, username string) (*domain.User, *domain.Article, error)
	ArticlePost(username string, article domain.Article) (*domain.User, *domain.Article, error)
	ArticlePut(username, slug string, fieldsToUpdate map[domain.ArticleUpdatableField]*string) (*domain.User, *domain.Article, error)
	ArticleDelete(username, slug string) error
}

type CommentsLogic interface {
	CommentsGet(slug string) ([]domain.Comment, error)
	CommentsPost(username, slug, comment string) (*domain.Comment, error)
	CommentsDelete(username, slug string, id int) error
}

type FavoritesLogic interface {
	FavoritesUpdate(username, slug string, favorite bool) (*domain.User, *domain.Article, error)
}

type TagsLogic interface {
	Tags(slug string) ([]string, error)
}

type HandlerConstructor struct {
}

func (c HandlerConstructor) New() Handler {
	return nil
}
