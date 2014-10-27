package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	r := render.New(render.Options{})
	router := mux.NewRouter()
	router.HandleFunc("/{name}/{age:[0-99]+}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		fmt.Fprintf(w, "Welcome "+vars["name"]+" you're "+vars["age"])
	})
	router.HandleFunc("/{name}/{age:[0-99]+}.json", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		r.JSON(w, http.StatusOK, map[string]string{"name": vars["name"], "age": vars["age"]})
	})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
