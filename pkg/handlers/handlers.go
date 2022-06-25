package handlers

import (
	"github.com/julienschmidt/httprouter"
	"evolve-credit/pkg/utils"
	"evolve-credit/pkg/repo"
	"log"
	"net/http"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]string)
	queryValues := r.URL.Query()

	params["from"] = queryValues.Get("from")
	params["to"] = queryValues.Get("to")
	params["limit"] = queryValues.Get("limit")
	params["offset"] = queryValues.Get("offset")

	users, err := repo.GetAll(params)
	if err != nil {
		log.Fatalln(err)
	}

	err = utils.WriteJSON(w, http.StatusOK, users, "users")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
	return
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	email := params.ByName("email")

	user, err := repo.Get(email)
	if err != nil {
		log.Fatalln(err)
	}

	err = utils.WriteJSON(w, http.StatusOK, user, "user")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}
