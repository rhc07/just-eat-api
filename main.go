package main

import (
	"encoding/json"
	"errors"
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
	vars := mux.Vars(r)
	postcode := vars["postcode"]
	err := Validate(postcode)
	if err != nil {
		log.Println("error: postcode is not valid")
		return
	}
	//For receiving API call
	client := http.Client{}
	apiRequest, httperr := http.NewRequest(http.MethodGet, APIURL+postcode, nil)
	if httperr != nil {
		log.Println("error: http request is not valid")
		return
	}
	apiResponse, geterr := client.Do(apiRequest)
	if geterr != nil {
		log.Println("error: no response from api")
		return
	}
	defer apiResponse.Body.Close()
	apiBody, readerr := ioutil.ReadAll(apiResponse.Body)
	if readerr != nil {
		log.Println("error: api body could not be read")
		return
	}
	restaurants := restaurantsList{}
	jsonerr := json.Unmarshal(apiBody, &restaurants)
	if jsonerr != nil {
		log.Println("error: json could not be unmarshalled")
		return
	}
	//For sending API call
	w.Header().Set("Access-Control-Allow-Origin", "*") //This heading is necessary for cross-origin data transfer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)

}

//Validate function checks the postcode from user input
func Validate(postcode string) error {
	if postcode == "" {
		return errors.New("invalid postcode: no input")
	}
	if len(postcode) > 7 {
		return errors.New("invalid postcode: too many characters")
	}
	if len(postcode) < 3 {
		return errors.New("invalid postcode: not enough characters")
	}
	return nil
}

func main() {
	//APIURL for API
	APIURL = "https://uk.api.just-eat.io/restaurants/bypostcode/"

	router := mux.NewRouter()
	router.HandleFunc("/postcode/{postcode}", ApiHandler)
	fmt.Println("Listening of port 8080")
	servErr := http.ListenAndServe(":8080", router)
	if servErr != nil {
		log.Fatalf("Server failed to start: %s", servErr)
	}
}
