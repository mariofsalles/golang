package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnFollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/followers",
		Method:       http.MethodGet,
		Function:     controllers.Followers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/following",
		Method:       http.MethodGet,
		Function:     controllers.Following,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/update-userpass",
		Method:       http.MethodPost,
		Function:     controllers.UpdatePassword,
		AuthRequired: true,
	},
}
