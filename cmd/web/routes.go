package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func wrapper(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	sessionChain := alice.New(app.session.Enable)

	router.GET("/", app.Index)
	router.GET("/drives/*stuff", app.Index)
	router.GET("/files/*foldername", app.FileList)

	router.POST("/files/*foldername", wrapper(sessionChain.Append(app.requireAuthentication).ThenFunc(app.FileUpload)))
	router.DELETE("/files/*foldername", wrapper(sessionChain.Append(app.requireAuthentication).ThenFunc(app.FileDelete)))
	router.POST("/folder/*foldername", wrapper(sessionChain.Append(app.requireAuthentication).ThenFunc(app.FolderUpload)))
	router.ServeFiles("/fileserver/*filepath", http.Dir(*app.globalPath))

	router.POST("/user/signup/"+*app.signUpSecret, wrapper(sessionChain.ThenFunc(app.signupUser)))
	router.POST("/user/login", wrapper(sessionChain.ThenFunc(app.loginUser)))
	router.POST("/user/logout", wrapper(sessionChain.Append(app.requireAuthentication).ThenFunc(app.logoutUser)))

	return alice.New(secureHeaders).Then(router)
}
