// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Tosaken"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/encounter": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "encounter"
                ],
                "summary": "すれ違いの取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thresher_usecase_model.Encounter"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/home": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "home"
                ],
                "summary": "家の作成",
                "parameters": [
                    {
                        "description": "Home",
                        "name": "Home",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InputHome"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thresher_usecase_model.Home"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/location": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "location"
                ],
                "summary": "現在地の作成",
                "parameters": [
                    {
                        "description": "Location",
                        "name": "Location",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InputLocation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thresher_usecase_model.Location"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "投稿の作成",
                "parameters": [
                    {
                        "description": "Description",
                        "name": "description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "投稿の取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "投稿の削除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/errors.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "ユーザーの更新",
                "parameters": [
                    {
                        "description": "User",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/thresher_usecase_model.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/follow/{id}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "新規フォロー",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/errors.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "フォロー解除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/errors.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/followed": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "フォローされてるユーザーの取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/thresher_usecase_model.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/following": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "フォロー中のユーザーの取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/thresher_usecase_model.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.ErrorResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "status_text": {
                    "type": "string"
                }
            }
        },
        "errors.SuccessResponse": {
            "type": "object",
            "properties": {
                "status_text": {
                    "type": "string"
                }
            }
        },
        "model.InputHome": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "non_pass_range": {
                    "type": "integer"
                }
            }
        },
        "model.InputLocation": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "model.Post": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "liked_num": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/thresher_usecase_model.Users"
                }
            }
        },
        "model.UpdateUser": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "thresher_usecase_model.Encounter": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "passed_id": {
                    "type": "string"
                },
                "passing_id": {
                    "type": "string"
                }
            }
        },
        "thresher_usecase_model.Home": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "non_pass_range": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/thresher_usecase_model.Users"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "thresher_usecase_model.Location": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "thresher_usecase_model.Users": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Thresher",
	Description:      "This is Thresher hackathon backend.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
