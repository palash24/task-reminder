package main

import (
	"fmt"
	"net/http"

	internal "github.com/palash24/task-reminder/internal"
)

func main() {
	internal.NewDb()
	server := internal.NewHttpServer()

	http.ListenAndServe(fmt.Sprintf("%s:%s", server.ServerAddr, server.ServerPort), server.Router)
}
