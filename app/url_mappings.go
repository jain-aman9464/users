package app

import(
	 "github.com/federicoleon/users/controllers/ping"
	 "github.com/federicoleon/users/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.FindUser)
	router.POST("/users", users.CreateUser)
}
