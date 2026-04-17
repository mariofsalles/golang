package routes

import "webapp/src/controllers"

var homePageRoute = Route{
	URI:          "/home",
	Method:       "GET",
	Function:     controllers.LoadHomePage,
	AuthRequired: true,
}