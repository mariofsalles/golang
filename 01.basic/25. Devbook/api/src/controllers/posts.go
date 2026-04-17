package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost creates a new post on the database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.PreparePost(); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.CreatePostOnDB(post)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts returns all posts from the database
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}
	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.GetPostsFromDB(userID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// GetPostById returns a post by its ID from the database
func GetPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post, err := repository.GetPostByIdFromDB(postID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, post)
}

// UpdatePost updates a post by its ID on the database
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postSavedInDB, err := repository.GetPostByIdFromDB(postID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedInDB.AuthorID != userID {
		responses.ErrJSON(w, http.StatusForbidden, errors.New("you can only update your own posts"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = post.PreparePost(); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePostOnDB(postID, post); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost deletes a post by its ID from the database
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postSavedInDB, err := repository.GetPostByIdFromDB(postID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedInDB.AuthorID != userID {
		responses.ErrJSON(w, http.StatusForbidden, errors.New("you can only delete your own posts"))
		return
	}

	if err = repository.DeletePostFromDB(postID); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetPostsByUserId returns all posts from a user by its ID from the database
func GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.GetPostsByUserIdFromDB(userID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// LikePost likes a post by its ID from the database
func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	if err = repository.LikePostOnDB(postID); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}


// UnlikePost unlikes a post by its ID from the database
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	if err = repository.UnlikePostOnDB(postID); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}