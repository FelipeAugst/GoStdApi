package routes

var Users = []Route{{
	Url:     "POST /users/create",
	Handler: nil,
},
	{Url: "GET /users",
		Handler: nil,
	},

	{Url: "GET /users/{id}",
		Handler: nil},
}
