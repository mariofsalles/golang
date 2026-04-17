package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User represents the user model using in the application
type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Userpass  string    `json:"userpass"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// GetFullUser calls the API 4 times to get all the user data
func GetFullUser(userID uint64, r *http.Request) (User, error) {
	userChan := make(chan User)
	followersChan := make(chan []User)
	followingsChan := make(chan []User)
	postsChan := make(chan []Post)

	go GetUserData(userChan, userID, r)
	go GetFollowers(followersChan, userID, r)
	go GetFollowing(followingsChan, userID, r)
	go GetPosts(postsChan, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChan:
			if userLoaded.ID == 0 {
				return User{}, errors.New("error fetching user data")
			}
			user = userLoaded
		case followersLoaded := <-followersChan:
			if followersLoaded == nil {
				return User{}, errors.New("error fetching followers data")
			}
			followers = followersLoaded
		case followingLoaded := <-followingsChan:
			if followingLoaded == nil {
				return User{}, errors.New("error fetching following data")
			}
			following = followingLoaded
		case postsLoaded := <-postsChan:
			if postsLoaded == nil {
				return User{}, errors.New("error fetching posts data")
			}
			posts = postsLoaded
		}
	}
	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// GetUserData fetch the users from the API
func GetUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf(`%s/users/%d`, config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// GetFollowers fetch the followers from the API
func GetFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf(`%s/users/%d/followers`, config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}
	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// GetFollowing fetch the following from the API
func GetFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf(`%s/users/%d/following`, config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followings []User
	if err = json.NewDecoder(response.Body).Decode(&followings); err != nil {
		channel <- nil
		return
	}
	if followings == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followings
}

// GetPosts fetch the posts from the API
func GetPosts(channel chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf(`%s/users/%d/posts`, config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	channel <- posts
}
