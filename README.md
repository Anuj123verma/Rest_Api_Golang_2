# State API

### An API that can fetch the **covid** related information about the state depending on the given **latitude** and **longitude**.

## Prerequisites

> Have an account on [here](https://developer.here.com/documentation/geocoding-search-api/dev_guide/topics/endpoint-geocode-brief.html).

## Instructions

> Generate the **API key** from the [**REST**](https://developer.here.com/projects/PROD-eba245db-dc9b-4902-9362-d66c48f39ad7) section on here platform.

## Running Instructions

* Parameters
    * DataType -- Format in which you want the output (String/JSON)
* QueryParameters
    * Latitude
    * Longitude
    * Database Name -- from which database you want to fetch the data (for me it is **Covid**)
    * Collection Name -- from whicn collection you want to fetch the data (for me it is **StateWise**)
    * ApiKey -- Key Generated from the [here](https://developer.here.com/documentation/geocoding-search-api/dev_guide/topics/endpoint-geocode-brief.html) platform
* Port
    * 8000 -- you can change it is **main.go**
* Type this on your browser to get the output
    * **http://localhost:8000/state/string?lat=""&long=""&db=""&collection=""&ApiKey=""**/ or 
    * **http://localhost:8000/state/json?lat=""&long=""&db=""&collection=""&ApiKey=""**/ 
    * Fill the enpty strings by yourself to get the desired output.
* For Swagger UI
    * **http://localhost:8000/swagger/index.html**

