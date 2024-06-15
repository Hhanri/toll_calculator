package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/Hhanri/toll_calculator/types"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "Listen Address")
	flag.Parse()

	store := NewMemoryStore()
	service := NewInvoiceAggregator(store)

	makeHttpTransport(*listenAddr, service)
}

func makeHttpTransport(listenAddr string, service Aggregator) {
	fmt.Println("Http transport running on", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(service))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(service Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var distance types.Distance
		err := json.NewDecoder(r.Body).Decode(&distance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
