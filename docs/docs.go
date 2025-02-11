// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/": {
            "delete": {
                "description": "deletes an image by user ID and checksum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "delete an image by user ID and checksum",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "checksum",
                        "name": "checksum",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            }
        },
        "/load": {
            "post": {
                "description": "create a image with specified parameter, image will be downloaded via source url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "create a image from external system",
                "parameters": [
                    {
                        "description": "body for upload a image",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ImageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            }
        },
        "/query": {
            "get": {
                "description": "Upload a image with specified parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "query image by external ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "externalID",
                        "name": "externalID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "Upload a image with specified parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "upload a image",
                "parameters": [
                    {
                        "description": "body for upload a image",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ImageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ImageRequest": {
            "type": "object",
            "required": [
                "algorithm",
                "externalComponent",
                "externalID",
                "fileName",
                "name",
                "userID"
            ],
            "properties": {
                "algorithm": {
                    "type": "string",
                    "enum": [
                        "md5",
                        "sha256"
                    ]
                },
                "checksum": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "externalComponent": {
                    "type": "string"
                },
                "externalID": {
                    "type": "string"
                },
                "fileName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "publish": {
                    "type": "boolean"
                },
                "sourceUrl": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Image": {
            "type": "object",
            "properties": {
                "algorithm": {
                    "type": "string"
                },
                "checksum": {
                    "type": "string"
                },
                "checksumPath": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "desc": {
                    "type": "string"
                },
                "externalComponent": {
                    "type": "string"
                },
                "externalID": {
                    "type": "string"
                },
                "fileName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imagePath": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "publish": {
                    "type": "boolean"
                },
                "sourceUrl": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "statusDetail": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
