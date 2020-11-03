package main

import(
	"fmt"
    "net/http"
    "github.com/irvandandung/goJWT/app"
    "github.com/irvandandung/goJWT/config"
)

func main(){
	mux := new(app.CustomMux)
	config.Routes(mux)

    server := new(http.Server)
    server.Handler = mux
    server.Addr = ":1234"

    fmt.Println("Starting server at", server.Addr)
    server.ListenAndServe()
}