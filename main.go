package main

import (
	"album-tracker/server"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/albuns/find", server.FindAlbum)
	http.HandleFunc("/albuns/register", server.RegisterAlbum)
	http.ListenAndServe(":8080", nil)
}
