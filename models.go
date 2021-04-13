package main

type OKResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type PostAuthzResponse struct {
	Message string `json:"message"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
