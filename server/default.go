package server

import (
	"bytes"
	"log"
	"net/http"

	"runtime/debug"

	"github.com/sahublr01/rest_crud_temp/utils"

	"github.com/sahublr01/rest_crud_temp/config"

	"github.com/kataras/golog"
)

var (
	BASE = []byte("/")
	API  = []byte("/api")
)
var OptionsAllowed = [][]byte{BASE}

func Init() {
	golog.Info("Starting Webserver on ", config.Props.Listen)
	handler := http.Handler(http.HandlerFunc(RequestHandler))
	if err := http.ListenAndServe(config.Props.Listen, handler); err != nil {
		log.Fatal(err)
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("%s\n%s", r, debug.Stack())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		}
	}()
	setCors(&w, r)
	if bytes.Equal([]byte(r.Method), []byte("OPTIONS")) {
		w.WriteHeader(http.StatusNoContent) //204
		return
	}
	HttpHandler(w, r)
}

func isCorsEnabled(r *http.Request) bool {
	for i := 0; i < len(OptionsAllowed); i++ {
		if utils.PathIs(OptionsAllowed[i], r) {
			return true
		}
	}
	return false
}

func setCors(w *http.ResponseWriter, r *http.Request) {
	if isCorsEnabled(r) {
		utils.SetCorsHeaders(w, r, true)
	}
}
