// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://google.com",
        "contact": {
            "name": "David Oheji",
            "url": "https://twitter.com/ejedavy",
            "email": "ejeohejidavid@gmail.com"
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
        "/api/v1/account/createaccount": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "description": "Create_Account",
                        "name": "CreateAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateAccountHandlerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/db.Account"
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/account/deleteaccount/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "We require all fields",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/account/getaccount/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/db.Account"
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/account/getaccounts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pageID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Account"
                            }
                        }
                    },
                    "400": {
                        "description": "We require all fields",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/account/updateaccount": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "description": "body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateAccountHandlerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/db.Account"
                        }
                    },
                    "400": {
                        "description": "We require all fields",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/log/getentry/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Entries"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/db.Entry"
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/log/listBankEntries": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Entries"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pageID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Entry"
                            }
                        }
                    },
                    "400": {
                        "description": "We require all fields",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/log/listaccountentries": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Entries"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pageID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "account_id",
                        "name": "account_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Entry"
                            }
                        }
                    },
                    "400": {
                        "description": "We require all fields",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/transfer/getTransfer/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/db.Transfer"
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/transfer/getincomingtransfers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pageID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "to_account_id",
                        "name": "to_account_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Transfer"
                            }
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/transfer/getoutgoingtransfers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "pageID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "from_account_id",
                        "name": "from_account_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Transfer"
                            }
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/transfer/initiatetransfer": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "parameters": [
                    {
                        "description": "requestBody",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.InitiateTransferHandlerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/db.TransferTXResult"
                        }
                    },
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/api.ServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CreateAccountHandlerRequest": {
            "type": "object",
            "required": [
                "currency",
                "owner"
            ],
            "properties": {
                "currency": {
                    "type": "string",
                    "enum": [
                        "USD",
                        "EUR",
                        "RUB",
                        "NGN"
                    ]
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "api.InitiateTransferHandlerRequest": {
            "type": "object",
            "required": [
                "amount",
                "receiver_id",
                "sender_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 1
                },
                "receiver_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "sender_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "api.ServerError": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "api.UpdateAccountHandlerRequest": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "db.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "db.Entry": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "description": "Can be positive or negative",
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "db.Transfer": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "This can only be positive",
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "to_account_id": {
                    "type": "integer"
                }
            }
        },
        "db.TransferTXResult": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "receiver": {
                    "$ref": "#/definitions/db.Account"
                },
                "receiverEntry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "sender": {
                    "$ref": "#/definitions/db.Account"
                },
                "senderEntry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "transfer": {
                    "$ref": "#/definitions/db.Transfer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Bank",
	Description:      "A simple bank API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
