package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ts, err := template.ParseFiles("./files/index.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func FileList(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	params := ps.ByName("foldername")
	dirList, err := fileDetailLister(*globalPath + "/drive" + params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dirList)
}
func FileUpload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := ps.ByName("foldername")
	newFilePath := *globalPath + "/drive" + params + "/"
	uploadFile(w, r, newFilePath)
}
func FolderUpload(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	params := ps.ByName("foldername")
	newFolderPath := *globalPath + "/drive" + params + "/"
	err := os.Mkdir(newFolderPath, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func FileDelete(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	params := ps.ByName("foldername")
	deletePathName := *globalPath + "/drive" + params
	err := os.Remove(deletePathName)
	if err != nil {
		errr := os.RemoveAll(deletePathName)
		if errr != nil {
			http.Error(w, errr.Error(), http.StatusInternalServerError)
		}
	}
	json.NewEncoder(w).Encode("Deleted")
}
