package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"syedzayyan.com/gosrvpi/pkg/forms"
	"syedzayyan.com/gosrvpi/pkg/models"
)

func (app *application) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (app *application) FileList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := ps.ByName("foldername")
	dirList, err := fileDetailLister(*app.globalPath + "/drive" + params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dirList)
}
func (app *application) FileUpload(w http.ResponseWriter, r *http.Request) {
	app.authState = app.isAuthenticated(r)
	ps := r.Context().Value("params").(httprouter.Params)
	params := ps.ByName("foldername")
	newFilePath := *app.globalPath + "/drive" + params + "/"
	uploadFile(w, r, newFilePath)
}
func (app *application) FolderUpload(w http.ResponseWriter, r *http.Request) {
	app.authState = app.isAuthenticated(r)
	ps := r.Context().Value("params").(httprouter.Params)
	params := ps.ByName("foldername")
	newFolderPath := *app.globalPath + "/drive" + params + "/"
	err := os.Mkdir(newFolderPath, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (app *application) FileDelete(w http.ResponseWriter, r *http.Request) {
	app.authState = app.isAuthenticated(r)
	ps := r.Context().Value("params").(httprouter.Params)
	params := ps.ByName("foldername")
	deletePathName := *app.globalPath + "/drive" + params
	err := os.Remove(deletePathName)
	if err != nil {
		errr := os.RemoveAll(deletePathName)
		if errr != nil {
			http.Error(w, errr.Error(), http.StatusInternalServerError)
		}
	}
	json.NewEncoder(w).Encode("Deleted")
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	app.authState = app.isAuthenticated(r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 8)

	// If there are any errors, redisplay the signup form.
	if !form.Valid() {
		fmt.Println(form.Errors)
		return
	}
	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Here rsdafo")
		}
		return
	}
	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	// And redirect the user to the login page.
	json.NewEncoder(w).Encode("Signed UP")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	app.authState = app.isAuthenticated(r)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	app.session.Put(r, "authenticatedUserID", id)

	// Redirect the user to the create snippet page.
	json.NewEncoder(w).Encode("Logged In")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.authState = app.isAuthenticated(r)
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	app.session.Remove(r, "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")
	json.NewEncoder(w).Encode("Logged Out")
}
