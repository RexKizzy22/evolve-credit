package main

import (
	// "errors"
	"fmt"
	// "github.com/julienschmidt/httprouter"
	// "log"
	"net/http"
	// "strconv"
)

func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {
	// queryValues := r.URL.Query()
	// numOfPages := queryValues.Get("pages")
	// fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("pages"))

	// users, err := app.models.DB.GetAll(numOfPages)
	// if err != nil {
	// 	app.logger.Fatalln(err)
	// }

	// err = app.writeJSON(w, http.StatusOK, users, "users")
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }
	fmt.Println("all users")
	return
}

// func (app *application) getUserById(w http.ResponseWriter, r *http.Request) {
// 	params := httprouter.ParamsFromContext(r.Context())
// 	queryValues := r.URL.Query()
// 	fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("name"))

// 	id, err := strconv.Atoi(params.ByName("id"))
// 	if err != nil {
// 		log.Println(errors.New("invalid id parameter"))
// 		return
// 	}

// 	movie, err := app.models.DB.Get(id)
// 	if err != nil {
// 		app.logger.Fatalln(err)
// 	}

// 	err = app.writeJSON(w, http.StatusOK, movie, "movie")
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// }
