package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/Hhanri/toll_calculator/types"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "Listen Address")
	flag.Parse()

	logger := NewLogrusLogger()
	store := NewMemoryStore()
	service := NewInvoiceAggregator(store)
	service = NewLogMiddleware(logger, service)

	makeHttpTransport(*listenAddr, service)
}

func makeHttpTransport(listenAddr string, service Aggregator) {
	fmt.Println("Http transport running on", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(service))
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		fmt.Printf("http listen and serve error: %s\n", err)
	}
}

func handleAggregate(service Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var distance types.Distance
		err := json.NewDecoder(r.Body).Decode(&distance)
		if err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}

		if err := service.AggregateDistance(distance); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, v error) error {
	return writeJson(w, status, map[string]string{"error": v.Error()})
}
