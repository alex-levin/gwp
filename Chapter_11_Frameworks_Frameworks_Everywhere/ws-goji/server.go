package main

import (
	"github.com/alex-levin/ws-goji/post"
	"github.com/zenazn/goji"

)

func main() {
	goji.Get("/post/:id", post.HandleGet)
	goji.Post("/post", post.HandlePost)
	goji.Put("/post/:id", post.HandlePut)
	goji.Delete("/post/:id", post.HandleDelete)
	goji.Serve()

}
