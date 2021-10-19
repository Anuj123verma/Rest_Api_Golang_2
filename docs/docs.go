// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Anuj Verma",
            "email": "anujssooni360@gmail.con"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/state/{data}": {
            "get": {
                "description": "Take Latitude, Longitude, ApiKey, Database name and Collection name as input. You can specify the data type (json/string)in which you want the response.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "States"
                ],
                "summary": "Give the covid related infomation about the state",
                "parameters": [
                    {
                        "type": "string",
                        "description": "datatype",
                        "name": "data",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "longitude",
                        "name": "long",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Database Name",
                        "name": "db",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Collection",
                        "name": "collection",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authentication Key",
                        "name": "ApiKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.State"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.Error"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "entity.State": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "integer"
                },
                "confirmed": {
                    "type": "integer"
                },
                "deaths": {
                    "type": "integer"
                },
                "recovered": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8000",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "Swagger Covid Information About State API",
	Description: "This is a server that can give the covid information about the states of India.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}