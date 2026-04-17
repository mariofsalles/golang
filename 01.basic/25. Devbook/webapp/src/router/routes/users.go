package routes

import (
	"webapp/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:          "/register-page",
		Method:       http.MethodGet,
		Function:     controllers.LoadRegisterUserPage,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsersPage,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.GetUserProfile,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Function:     controllers.LoadingCurrentUserProfilePage,
		AuthRequired: true,
	},
	{
		URI:          "/update-profile",
		Method:       http.MethodGet,
		Function:     controllers.LoadingUpdateCurrentUserPage,
		AuthRequired: true,
	},
	{
		URI:          "/update-profile",
		Method:       http.MethodPut,
		Function:     controllers.UpdateCurrentUser,
		AuthRequired: true,
	},
	{
		URI:          "/update-password",
		Method:       http.MethodGet,
		Function:     controllers.LoadingUpdatePasswordPage,
		AuthRequired: true,
	},
	{
		URI:          "/update-password",
		Method:       http.MethodPost,
		Function:     controllers.UpdatePassword,
		AuthRequired: true,
	},
	{
		URI:          "/delete-user",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
}
