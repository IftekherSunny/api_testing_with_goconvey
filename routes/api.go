package routes

import (
	"github.com/iftekhersunny/api_testing_with_goconvey/controllers"
	"github.com/iftekhersunny/api_testing_with_goconvey/router"
)

// API version number constant value
const v1 = "/v1"

// Controller instance
var userController = new(controllers.UserController)

var ApiRoutes = router.Routes{

	// user controller routes...
	router.Route{
		Path:    "/users",
		Method:  "GET",
		Handler: userController.Index,
		Version: v1,
	},

	router.Route{
		Path:    "/users",
		Method:  "POST",
		Handler: userController.Create,
		Version: v1,
	},
}
