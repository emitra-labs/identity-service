package model

type HelloRequest struct {
	Name string `params:"name" path:"name"`
}

type HelloResponse struct {
	Message string
}
