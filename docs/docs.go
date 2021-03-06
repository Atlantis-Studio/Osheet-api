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
        "license": {
            "name": "MIT",
            "url": "https://www.mit.edu/~amini/LICENSE.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/channels": {
            "get": {
                "description": "Get channels by filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Channels"
                ],
                "summary": "Get Channels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Channel"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Store new channel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Channels"
                ],
                "summary": "Create Channl",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Display Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Twitter Account",
                        "name": "twitterAccount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Avatar Photo Url",
                        "name": "avatar",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Company",
                        "name": "company",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Unit",
                        "name": "unit",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Channel Url",
                        "name": "channelUrl",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Birthday",
                        "name": "birthday",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Height",
                        "name": "height",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Debut Date",
                        "name": "debutDate",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Channel"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/channels/{twitterAccount}": {
            "get": {
                "description": "Show channel info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Channels"
                ],
                "summary": "Get Channel By Twitter Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Twitter Account",
                        "name": "twitterAccount",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Channel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Channel": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "channelUrl": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "debutDate": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "height": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "twitterAccount": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                },
                "updatedAt": {
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Oshi-api",
	Description: "Oshi-api swagger",
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
	swag.Register("swagger", &s{})
}
