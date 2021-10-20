package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Anuj123Verma/Rest_Api_Golang_2/entity"
	"github.com/Anuj123Verma/Rest_Api_Golang_2/mongodb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateResource godoc
// @Summary Give the covid related infomation about the state
// @Description Take Latitude, Longitude, ApiKey, Database name and Collection name as input. You can specify the data type (json/string)in which you want the response.
// @Tags States
// @Param dataType path string true "datatype"
// @Param lat query string true "latitude"
// @Param long query string true "longitude"
// @Param db query string true "Database Name"
// @Param collection query string true "Collection"
// @Param ApiKey query string true "Authentication Key"
// @Consume  application/json
// @Produce json
// @Success 200  {object} entity.State
// @Failure 400  {object} entity.Error
// @Failure 500  {string} string "error"
// @Router /state/{data} [get]
func Getstate(c echo.Context) error {
	// stroing the parameters and querparameters in variables
	dataType := c.Param("dataType")
	latitude := c.QueryParam("lat")
	longitude := c.QueryParam("long")
	dbname := c.QueryParam("db")
	collectionname := c.QueryParam("collection")
	key := c.QueryParam("ApiKey")

	// getting the link of the api from which fetching the covid information
	link := Getlink(latitude, longitude, key)
	// Geo struct type variable declared
	var locations entity.Geo
	// consuming the api to get the location for a particular latitude and longitude
	var nlocations entity.Geo = ConsumeApi(link, locations)
	// checking if the location lies in India
	if nlocations.Items[0].Address.CountryName == "India" {
		// state struct type variable declared
		var state entity.State
		// creating a filter on the basis of state name
		filter := bson.D{primitive.E{Key: "state", Value: nlocations.Items[0].Address.State}}
		// fetching the collection from mongodb
		collection := mongodb.Getdb(dbname, collectionname)
		// fetching the information from the collection based on filter
		err := collection.FindOne(context.Background(), filter).Decode(&state)
		if err != nil {
			log.Fatal(err)
		}
		// dataType checked
		if dataType == "string" {
			return c.String(http.StatusOK, fmt.Sprintf("state : %s\nconfirmed cases : %d\nrecovered cases : %d\ntotal deaths : %d\ntotal active cases : %d\n", state.State, state.Confirmed, state.Recovered, state.Deaths, state.Active))
		}
		if dataType == "json" {
			return c.JSON(http.StatusOK, &entity.State{
				State:     state.State,
				Confirmed: state.Confirmed,
				Recovered: state.Recovered,
				Deaths:    state.Deaths,
				Active:    state.Active,
			})
		}
		// dataType is not json and string
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "give the correct datatype",
		})

	}
	// if the latitude and longitude does not belongs to India
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "give the correct latitude and longitude",
	})

}
