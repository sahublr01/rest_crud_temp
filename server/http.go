package server

import (
	"fmt"
	"net/http"

	// "rest_crud/app"
	"github.com/sahublr01/rest_crud_temp/utils"

	"github.com/sahublr01/rest_crud_temp/app"

	"github.com/kataras/golog"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if utils.PathIs(BASE, r) {
		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}
	} else if utils.PathIs(API, r) {
		if r.Method == "POST" {
			golog.Info("Post method check")
			fmt.Print("Post method check")
			res := app.PostAPI(w, r)
			utils.SendJSONResponse(w, res, "Success", "", http.StatusOK, http.StatusOK)
			return
		} else if r.Method == "GET" {
			golog.Info("Get Method check")
			res := app.GetAPI(w, r)
			utils.SendJSONResponse(w, res, "Success", "", http.StatusOK, http.StatusOK)
		} else if r.Method == "PUT" {
			golog.Info("Put Method Check")
			res := app.PutAPI(w, r)
			utils.SendJSONResponse(w, res, "Success", "", http.StatusOK, http.StatusOK)
			return
		} else if r.Method == "DELETE" {
			golog.Info("Delete Method Check")
			res := app.DeleteAPI(w, r)
			utils.SendJSONResponse(w, res, "Success", "", http.StatusOK, http.StatusOK)
			return
		} else {
			utils.SendJSONResponse(w, nil, "Success", "", http.StatusBadRequest, http.StatusBadRequest)
		}
	} else {
		utils.SendJSONResponse(w, nil, "page not found", "", http.StatusNotFound, http.StatusNotFound)
	}
}
