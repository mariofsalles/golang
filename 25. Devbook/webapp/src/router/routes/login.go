package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:          "/",
		Function:     controllers.LoadLoginPage,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Function:     controllers.LoadLoginPage,
		Method:       http.MethodGet,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Function:     controllers.LoginUser,
		Method:       http.MethodPost,
		AuthRequired: false,
	},
}
