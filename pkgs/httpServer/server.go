package httpServer

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func newMuxRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		
	})

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa).Methods("GET")
	return r
}

func RunServer(port int) {
	router := newMuxRouter()
	portNum := strconv.Itoa(port)
	srv := &http.Server{
		Handler: router,
		Addr:    ":"+portNum,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
