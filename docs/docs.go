// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ticket": {
            "get": {
                "description": "GET TICKETS",
                "tags": [
                    "GET TICKETS"
                ],
                "summary": "GET TICKETS",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Start date (Unix time)",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "End date (Unix time)",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "PENDING",
                            "ACCEPTED",
                            "RESOLVED",
                            "REJECTED"
                        ],
                        "type": "string",
                        "description": "Filter by ticket status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page of pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "minimum": 100,
                        "type": "integer",
                        "description": "total record to show",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Ticket"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "UPDATE TICKET",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UPDATE TICKETS"
                ],
                "summary": "UPDATE TICKET",
                "parameters": [
                    {
                        "description": "Body of ticket",
                        "name": "ticketInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TicketStatusUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update ticket successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "CREATE TICKET",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CREATE TICKETS"
                ],
                "summary": "CREATE TICKET",
                "parameters": [
                    {
                        "description": "Body of ticket",
                        "name": "ticketInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TicketCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Ticket"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.Error": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "model.Ticket": {
            "type": "object",
            "properties": {
                "contactInfo": {
                    "type": "string"
                },
                "createAt": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                }
            }
        },
        "model.TicketCreateRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "contactInfo": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.TicketStatusUpdateRequest": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "string"
                }
            }
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
	Host:        "127.0.0.1:1323",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
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
