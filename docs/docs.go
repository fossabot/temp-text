// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Source Code",
            "url": "https://github.com/sixwaaaay/temp-text"
        },
        "license": {
            "name": "Apache 2.0 License",
            "url": "https://github.com/sixwaaaay/temp-text/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/query": {
            "get": {
                "description": "query the text by tid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Query",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tid",
                        "name": "tid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/logic.Resp-string"
                        }
                    }
                }
            }
        },
        "/share": {
            "post": {
                "description": "share the text",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Share",
                "parameters": [
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/logic.Resp-string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "logic.Resp-string": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Temp-text API",
	Description:      "temporary text storage",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}