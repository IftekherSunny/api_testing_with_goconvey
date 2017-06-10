package routes

import (
	"github.com/iftekhersunny/api_testing_with_goconvey/controllers"
	"github.com/iftekhersunny/api_testing_with_goconvey/router"
)

// API version number constant value
const v1 = "/v1"

// Controller instance
var usersController = new(controllers.UsersController)

var ApiRoutes = router.Routes{

	// users controller routes...
	router.Route{
		Path:    "/users",
		Method:  "GET",
		Handler: usersController.Index,
		Version: v1,
	},

	router.Route{
		Path:    "/users",
		Method:  "POST",
		Handler: usersController.Create,
		Version: v1,
	},
}
