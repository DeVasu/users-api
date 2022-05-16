package app

import (
	"bookstore/token-jwt/controllers/ping"
	"bookstore/token-jwt/controllers/tokens"
	"bookstore/token-jwt/controllers/users"
)

func mapurls() {
	// to check if the server is working
	router.GET("/ping", ping.Ping)

	//for signing up with name, email, password
	router.POST("/signup", users.CreateUser)

	//for loggin in user with email, password, returns a token
	router.POST("/login", users.Login)

	//to login by token only, returns userId to the caller
	router.GET("/token/login", tokens.GetUserByToken)

	// USERS handlers
	// router.GET("/internal/users/search", users.Search)
	// router.GET("/users/:user_id", users.GetUser)
	// router.POST("/users", users.CreateUser)
	// router.PUT("/users/:user_id", users.UpdateUser)
	// router.PATCH("/users/:user_id", users.UpdateUser)
	// router.DELETE("/users/:user_id", users.DeleteUser)
	// router.POST("/users/login", users.Login)
	// router.GET("/token/:token_id", users.TokenLogin)
}
