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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/members": {
            "get": {
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "summary": "List members",
                "operationId": "list-members",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Go to the page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "How many member per pages",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "Create a new member, employee only accepts role, contractor only accepts contract_duration",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "summary": "Create a new member",
                "operationId": "create-member",
                "parameters": [
                    {
                        "description": "Member Payload",
                        "name": "member",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/validate.Errors"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/members/{member_id}": {
            "get": {
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "summary": "Show a member",
                "operationId": "show-member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Member ID",
                        "name": "member_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "summary": "Update a member",
                "operationId": "update-member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Member ID",
                        "name": "member_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/validate.Errors"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "summary": "Delete a member",
                "operationId": "delete-member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Member ID",
                        "name": "member_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Member": {
            "type": "object",
            "properties": {
                "contract_duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "employee",
                        "contractor"
                    ]
                }
            }
        },
        "validate.Errors": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
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
	Host:        "localhost:3000",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Team Manager API",
	Description: "RESTful API that will help you manage your team. You can create a member (employee or contractor) and attach a tag to him.",
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
