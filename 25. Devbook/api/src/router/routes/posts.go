package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoute = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Function:     controllers.GetPosts,
		AuthRequired: false,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodGet,
		Function:     controllers.GetPostById,
		AuthRequired: false,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodPost,
		Function:     controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		AuthRequired: true,
	},
	{
		URI:  "/users/{userID}/posts",
		Method:       http.MethodGet,
		Function:     controllers.GetPostsByUserId,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePost,
		AuthRequired: true,
	},
}
