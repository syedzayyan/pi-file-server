package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type files struct {
	FileName     string `json:"filename"`
	FileSize     int64  `json:"filesize"`
	FileModTime  string `json:"filemodtime"`
	FileOrFolder bool   `json:"fileorfolder"`
}

func fileDetailLister(root string) ([]files, error) {
	c, err := ioutil.ReadDir(root)

	fileMap := []files{}
	for _, entry := range c {
		tempStruct := files{entry.Name(), entry.Size(), entry.ModTime().Format("15:04:05 || Jan 2 2006"), entry.IsDir()}
		fileMap = append(fileMap, tempStruct)
	}

	return fileMap, err
}
func uploadFile(w http.ResponseWriter, r *http.Request, dirListUp string) {

	r.ParseMultipartForm(32)
	multiFormFields := r.MultipartForm
	files := multiFormFields.File["files"]
	for index := range files {
		f, err := files[index].Open()
		if err != nil {
			f.Close()
			panic(err)
		}
		defer f.Close()

		dst, err := os.Create(dirListUp + files[index].Filename)
		if err != nil {
			fmt.Println("Error Retrieving the File")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			dst.Close()
			return
		}
		defer dst.Close()
		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, f); err != nil {
			fmt.Println("Error Retrieving the File")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	multiFormFields.RemoveAll()
	w.Write([]byte("{'data':'Successfully Uploaded File'}"))
}
