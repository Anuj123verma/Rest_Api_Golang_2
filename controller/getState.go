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
	dataType := c.Param("dataType")
	latitude := c.QueryParam("lat")
	longitude := c.QueryParam("long")
	dbname := c.QueryParam("db")
	collectionname := c.QueryParam("collection")
	key := c.QueryParam("ApiKey")

	link := Getlink(latitude, longitude, key)
	var locations entity.Geo
	var nlocations entity.Geo = ConsumeApi(link, locations)
	if nlocations.Items[0].Address.CountryName == "India" {
		var state entity.State
		filter := bson.D{primitive.E{Key: "state", Value: nlocations.Items[0].Address.State}}
		collection := mongodb.Getdb(dbname, collectionname)
		err := collection.FindOne(context.Background(), filter).Decode(&state)
		if err != nil {
			log.Fatal(err)
		}
		if dataType == "string" {
			return c.String(http.StatusOK, fmt.Sprintf("state : %s\nconfirmed cases : %d\nrecovered cases : %d\ntotal deaths : %d\ntotal active cases : %d\n", state.State, state.Confirmed, state.Recovered, state.Deaths, state.Active))
		}
		if dataType == "json" {
			return c.JSON(http.StatusOK, map[string]string{
				"state":              state.State,
				"confirmed cases":    strconv.Itoa(state.Confirmed),
				"recovered cases":    strconv.Itoa(state.Recovered),
				"total deaths":       strconv.Itoa(state.Deaths),
				"total active cases": strconv.Itoa(state.Active),
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "give the correct datatype",
		})

	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "give the correct latitude and longitude",
	})

}
