package main

import (
	"encoding/json"
	"mock-coupon-api/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jaswdr/faker"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)

var Faker = faker.New()

func updateOffersLoop() {
	updateDelayRaw := os.Getenv("UPDATE_DELAY_SECONDS")

	updateDelay, _ := strconv.Atoi(updateDelayRaw)

	for {

		funcs := []func(){
			database.AddNewFakeOffer,
			database.UpdateOffer,
			database.SuspendOffer,
		}

		choice := Faker.IntBetween(0, 2)

		funcs[choice]()

		time.Sleep(time.Duration(time.Second * time.Duration(updateDelay)))

	}
}

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {

	rawLastExtract := r.URL.Query().Get("last_extract")

	if rawLastExtract == "" {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")

		response := OfferResponse{
			Result: true,
			Offers: database.GetOffers(),
		}

		data, _ := json.Marshal(response)
		w.Write(data)
		return
	}

	lastExtract, err := strconv.Atoi(rawLastExtract)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lastFetchedTime := time.Unix(int64(lastExtract), 0)
	offers := database.FetchOfferAfter(lastFetchedTime)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	response := OfferResponse{
		Result: true,
		Offers: offers,
	}
	data, _ := json.Marshal(response)
	w.Write(data)
	return

}

func main() {

	server := http.NewServeMux()

	for i := 0; i < 5; i++ {
		database.AddNewFakeOffer()
	}

	go updateOffersLoop()

	server.HandleFunc("/", httpRequestHandler)
	server.Handle("/metrics", promhttp.Handler())

	// http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", server)

}