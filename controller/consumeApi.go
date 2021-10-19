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
	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &locations)
	return locations
}
