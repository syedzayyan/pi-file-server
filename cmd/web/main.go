package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

var globalPath = flag.String("hostDir", "./files", "HTTP network address")

func main() {
	addr := flag.String("addr", ":5000", "HTTP network address")
	flag.Parse()

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/drives/*stuff", Index)

	router.GET("/files/*foldername", FileList)
	router.POST("/files/*foldername", FileUpload)
	router.DELETE("/files/*foldername", FileDelete)

	router.POST("/folder/*foldername", FolderUpload)

	router.ServeFiles("/fileserver/*filepath", http.Dir(*globalPath))

	aliceMainChain := alice.New(secureHeaders).Then(router)

	log.Fatal(http.ListenAndServe(*addr, aliceMainChain))
}
