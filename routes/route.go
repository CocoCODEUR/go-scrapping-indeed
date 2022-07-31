package routes

import (
	"log"

	"github.com/CocoCODEUR/go-scrapping-indeed/scrapping"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getJobsData(w *http.ResponseWriter, r *http.Request) {

	data := scrapping.Scrapper()
	w.Header().Set("Content-type", "applications/json")
	json.NewEncoder(w).Encode(data)

}

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/jobData", getJobsData).Methods("GET")

	fmt.Println("status server at  port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
