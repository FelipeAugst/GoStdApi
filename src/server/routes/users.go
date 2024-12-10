package routes

import "api/src/controllers"

var Users = []Route{{
	Url:     "/users/create",
	Method:  "POST",
	Handler: controllers.CreateUser,
},
	{Url: "/users/",
		Method:  "GET",
		Handler: controllers.GetUsers,
	},

	{Url: "/users/{id}",
		Method:  "GET",
		Handler: controllers.FindUser},
	{Url: "/users/{id}/update",
		Method:  "PUT",
		Handler: controllers.UpdateUser},
	{Url: "/users/{id}/delete",
		Method:  "DELETE",
		Handler: controllers.DeleteUser},
}
