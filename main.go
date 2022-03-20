package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//APIURL variable is set in the main function
var (
	APIURL string
)

//restaurants struct sets up the required properties of each restaurant
type restaurant struct {
	Name          string  `json:"name"`
	RatingAverage float32 `json:"ratingAverage"`
}

//restaurantsList is a list of each restaurant
type restaurantsList struct {
	Restaurants []restaurant `json:"restaurants"`
}

//ApiHandler sends a http Get request to the api to fetch the data
func ApiHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Access-Control-Allow-Origin", "*") //This heading is necessary for cross-origin data transfer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)

}

func GetInput() string {
	fmt.Println("Enter your postcode (please do not leave any whitespace between characters): ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	return input
}

//GetPostcode function gets a postcode from user input
func GetPostcode(input string) string {
	if input == "" {
		log.Fatalf("Invalid postcode: no input")
	}
	if len(input) > 7 {
		log.Fatalf("Invalid postcode: too many characters")
	}
	if len(input) < 5 {
		log.Fatalf("Invalid postcode: not enough characters")
	}
	return input
}

func main() {
	input := GetInput()
	postcode := GetPostcode(input)

	//APIURL for API
	APIURL = "https://uk.api.just-eat.io/restaurants/bypostcode/" + postcode

	router := mux.NewRouter()
	router.HandleFunc("/", ApiHandler)
	fmt.Println("Listening of port 8080")
	servErr := http.ListenAndServe(":8080", router)
	if servErr != nil {
		log.Fatalf("Server failed to start: %s", servErr)
	}
}
