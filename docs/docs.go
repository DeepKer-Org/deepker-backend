// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/tests": {
            "get": {
                "description": "Get all tests",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "Get all tests",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Test"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new test",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "Create a new test",
                "parameters": [
                    {
                        "description": "Test",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Test"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Test"
                        }
                    }
                }
            }
        },
        "/tests/{id}": {
            "get": {
                "description": "Get a single test by DoctorID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "Get test by DoctorID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Test DoctorID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Test"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing test",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "Update a test",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Test DoctorID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Test",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Test"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Test"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a test by DoctorID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "Delete a test",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Test DoctorID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a single user by DoctorID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by DoctorID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User DoctorID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User DoctorID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by DoctorID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User DoctorID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Test": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
