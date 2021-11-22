package domain

import (
	"sort"
	"time"
)

// User represents a user account in the system
type User struct {
	Name      string
	Email     string
	Password  string
	Bio       *string
	ImageLink *string
	FollowIDs []string
	Favorites []Article
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdatableProperty int

const (
	UserEmail UserUpdatableProperty = iota
	UserName
	UserBio
	UserImageLink
	UserPassword
)

func UpdateUser(initial *User, opts ...func(*User)) {
	for _, v := range opts {
		v(initial)
	}
}

func SetUserName(input *string) func(*User) {
	return func(user *User) {
		if input != nil {
			user.Name = *input
		}
	}
}

func SetUserEmail(input *string) func(*User) {
	return func(user *User) {
		if input != nil {
			user.Email = *input
		}
	}
}

func SetUserBio(input *string) func(*User) {
	return func(user *User) {
		if input == nil {
			return
		}
		if *input == "" {
			user.Bio = nil
		} else {
			user.Bio = input
		}
	}
}

func SetUserImageLink(input *string) func(*User) {
	return func(user *User) {
		if input == nil {
			return
		}
		if *input == "" {
			user.ImageLink = nil
		} else {
			user.ImageLink = input
		}
	}
}

func SetUserPassword(input *string) func(*User) {
	return func(user *User) {
		if input != nil {
			user.Password = *input
		}
	}
}

func (user *User) Follows(userName string) bool {
	if user.FollowIDs == nil {
		return false
	}
	sort.Strings(user.FollowIDs)
	i := sort.SearchStrings(user.FollowIDs, userName)
	return i < len(user.FollowIDs) && user.FollowIDs[i] == userName
}

func (user *User) UpdateFollowees(userName string, follow bool) {
	if follow && !user.Follows(userName) {
		user.FollowIDs = append(user.FollowIDs, userName)
	} else {
		for i, followee := range user.FollowIDs {
			if followee == userName {
				user.FollowIDs = append(user.FollowIDs[0:i], user.FollowIDs[i+1:len(user.FollowIDs)]...)
			}
		}
	}
	if len(user.FollowIDs) == 0 {
		user.FollowIDs = nil
	}
}
