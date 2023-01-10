package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:              "/users",
		Method:           http.MethodPost,
		Function:         controllers.CreateUser,
		MustAuthenticate: false,
	},
	{
		URI:              "/users",
		Method:           http.MethodGet,
		Function:         controllers.GetAllUsers,
		MustAuthenticate: false,
	},
	{
		URI:              "/users/{id}",
		Method:           http.MethodGet,
		Function:         controllers.FindUser,
		MustAuthenticate: false,
	},
	{
		URI:              "/users/{id}",
		Method:           http.MethodPut,
		Function:         controllers.UpdateUser,
		MustAuthenticate: false,
	},
	{
		URI:              "/users/{id}",
		Method:           http.MethodDelete,
		Function:         controllers.DeleteUser,
		MustAuthenticate: false,
	},
}
