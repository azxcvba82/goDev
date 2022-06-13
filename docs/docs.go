// Package docs GENERATED BY SWAG; DO NOT EDIT
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
        "/allAlbumType": {
            "get": {
                "description": "allAlbumType load",
                "tags": [
                    "Search"
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/allkind": {
            "get": {
                "description": "allkind load",
                "tags": [
                    "Kind"
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/api/getAccountFromJWT": {
            "get": {
                "security": [
                    {
                        "securityDefinitions.apikey BearerAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "getAccountFromJWT",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/api/getShoppingCartByAccount": {
            "get": {
                "security": [
                    {
                        "securityDefinitions.apikey BearerAuth": []
                    },
                    {
                        "BearerAuth": []

                    }
                ],
                "description": "getShoppingCartByAccount load",
                "tags": [
                    "ShoppingCart"
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/getAlbumById": {
            "get": {
                "description": "getAlbumById load",
                "tags": [
                    "Album"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/getAlbumsByKindId": {
            "get": {
                "description": "getAlbumsByKindId load",
                "tags": [
                    "Album"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "kindId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/getProductsByAlbumId": {
            "get": {
                "description": "getProductsByAlbumId load",
                "tags": [
                    "Album"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "albumId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/getProductsByProductName": {
            "get": {
                "description": "getProductsByProductName load",
                "tags": [
                    "Search"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "parameters": [
                    {
                        "description": "json",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/mainActivities": {
            "get": {
                "description": "mainActivities load",
                "tags": [
                    "Homepage"
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/mainAlbums": {
            "get": {
                "description": "mainAlbums load",
                "tags": [
                    "Homepage"
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "parameters": [
                    {
                        "description": "json",
                        "name": "signup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "This is a api web server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
