package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	APIKEY   string
	APIURL   string
	POSTCODE string
)

type restaurant struct {
	Name          string  `json:"name"`
	RatingAverage float32 `json:"ratingAverage"`
}

type restaurantsList struct {
	Restaurants []restaurant `json:"restaurants"`
}

func apiCall(w http.ResponseWriter, r *http.Request) {
	//For receiving API call
	client := http.Client{}
	apiRequest, httperr := http.NewRequest(http.MethodGet, APIURL, nil)
	if httperr != nil {
		log.Fatal(httperr)
	}
	apiResponse, geterr := client.Do(apiRequest)
	if geterr != nil {
		log.Fatal(geterr)
	}
	apiBody, readerr := ioutil.ReadAll(apiResponse.Body)
	if readerr != nil {
		log.Fatal(readerr)
	}
	restaurants := restaurantsList{}
	jsonerr := json.Unmarshal(apiBody, &restaurants)
	if jsonerr != nil {
		log.Fatal(jsonerr)
	}
	//For sending API call
	w.Header().Set("Access-Control-Allow-Origin", "*") //This heading is necessary for cross origin data transfer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)

}

func main() {
	fmt.Println("Enter your postcode: ")
	var postcode string
	fmt.Scanln(&postcode)
	//APIURL for API
	APIURL = "https://uk.api.just-eat.io/restaurants/bypostcode/" + postcode

	router := mux.NewRouter()
	//router.HandleFunc("/postcode", getPostcodeHandler)
	router.HandleFunc("/", apiCall)
	fmt.Println("Listening of port 8080")
	http.ListenAndServe(":8080", router)
}
