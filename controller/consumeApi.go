package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Anuj123Verma/Rest_Api_Golang_2/entity"
)

func ConsumeApi(link string, locations entity.Geo) entity.Geo {
	// getting the response from the link
	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// getting the responseData
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Unmarshalling with json
	json.Unmarshal(responseData, &locations)

	// return the Geo entity and store the data in that
	return locations
}
