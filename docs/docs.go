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
        "/coin/list": {
            "get": {
                "description": "get coinlist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coin"
                ],
                "summary": "Coin-List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Resp_Quote"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/coin/quote": {
            "patch": {
                "description": "get coinquote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coin"
                ],
                "summary": "websocket",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Resp_Quote"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/coin/wallet": {
            "get": {
                "description": "get coinwallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coin"
                ],
                "summary": "Coin-Wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Wallet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "로그인",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "operationId": "post-user-signin",
                "parameters": [
                    {
                        "description": "User",
                        "name": "emailPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.emailPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "회원가입",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register",
                "operationId": "post-user-signup",
                "parameters": [
                    {
                        "description": "User",
                        "name": "model.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.emailPassword": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "jjh123@naver.com"
                },
                "password": {
                    "type": "string",
                    "example": "123"
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.Resp_Quote": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "required": [
                            "name",
                            "symbol"
                        ],
                        "properties": {
                            "name": {
                                "type": "string"
                            },
                            "quote": {
                                "type": "object",
                                "properties": {
                                    "BTC": {
                                        "type": "object",
                                        "properties": {
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    },
                                    "USD": {
                                        "type": "object",
                                        "properties": {
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                }
                            },
                            "symbol": {
                                "type": "string"
                            }
                        }
                    }
                },
                "status": {
                    "type": "object",
                    "properties": {
                        "timestamp": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "access_key",
                "email",
                "name",
                "password",
                "phone_number",
                "secret_key"
            ],
            "properties": {
                "access_key": {
                    "type": "string",
                    "example": "cY158XlCRODQljHva8pMjORsoxrKRdfg4S9jT8qa"
                },
                "email": {
                    "type": "string",
                    "example": "jjh123@naver.com"
                },
                "name": {
                    "type": "string",
                    "example": "jjh"
                },
                "password": {
                    "type": "string",
                    "example": "123"
                },
                "phone_number": {
                    "type": "string",
                    "example": "010-1234-5678"
                },
                "secret_key": {
                    "type": "string",
                    "example": "2y0BcdVYH48Hxc8SEwfOucxAqMoL623K70j6OCWa"
                }
            }
        },
        "model.Wallet": {
            "type": "object",
            "properties": {
                "avg_buy_price": {
                    "type": "string"
                },
                "avg_buy_price_modified": {
                    "type": "boolean"
                },
                "balance": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "locked": {
                    "type": "string"
                },
                "unit_currency": {
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
	Host:        "localhost:5000",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server celler server.",
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
