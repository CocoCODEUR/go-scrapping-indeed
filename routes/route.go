package routes

import (
	"log"

	"github.com/CocoCODEUR/go-scrapping-indeed/scrapping"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getJobsData(w http.ResponseWriter, r *http.Request) {
	//CORS Enable
	enableCors(&w)

	data := scrapping.Scrapper()
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func CreateRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/jobData", getJobsData).Methods("GET")

	fmt.Printf("status server at  port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
