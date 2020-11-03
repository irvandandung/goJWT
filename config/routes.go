package config

import(
	"github.com/irvandandung/goJWT/app"
)

func Routes(mux *app.CustomMux){
	mux.HandleFunc("/login", app.HandlerLogin)
	mux.RegisterMiddleware(app.MiddlewareJWTAuthorization)
    mux.HandleFunc("/index", app.HandlerIndex)
}