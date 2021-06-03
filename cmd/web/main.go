package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/drives/*stuff", Index)

	router.GET("/files/*foldername", FileList)
	router.POST("/files/*foldername", FileUpload)
	router.DELETE("/files/*foldername", FileDelete)

	router.POST("/folder/*foldername", FolderUpload)

	router.ServeFiles("/fileserver/*filepath", http.Dir("./files"))

	aliceMainChain := alice.New(secureHeaders).Then(router)

	log.Fatal(http.ListenAndServe(":8080", aliceMainChain))
}
