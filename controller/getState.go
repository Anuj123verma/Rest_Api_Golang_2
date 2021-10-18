package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Anuj123Verma/Rest_Api_Golang_2/entity"
	"github.com/Anuj123Verma/Rest_Api_Golang_2/mongodb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Getstate(c echo.Context) error {
	latitude := c.QueryParam("lat")
	longitude := c.QueryParam("long")
	dbname := c.QueryParam("db")
	collectionname := c.QueryParam("collection")
	key := c.QueryParam("key")
	dataType := c.Param("data")

	link := Getlink(latitude, longitude, key)
	var locations entity.Geo
	var nlocations entity.Geo = ConsumeApi(link, locations)

	if nlocations.Items[0].Address.CountryName == "India" {
		var state entity.State

		filter := bson.D{primitive.E{Key: "state", Value: nlocations.Items[0].Address.State}}
		collection := mongodb.Getdb(dbname, collectionname)
		err := collection.FindOne(context.Background(), filter).Decode(&state)
		if err != nil {
			fmt.Println(err)
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

	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "give the correct latitude and longitude",
	})

}
