package main

import (
    "log"
    "net/http"

    handlers "main/handlers"
    "github.com/rs/cors"
)

func main() {
    router := handlers.NewRouter()

    // CORS middleware
    c := cors.AllowAll()

    // Handler with CORS support
    handler := c.Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}
