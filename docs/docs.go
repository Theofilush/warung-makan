package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
	"swagger": "2.0",
	"info": {
	  "version": "1.0",
	  "title": "Warung Makan",
	  "contact": {}
	},
	"host": "localhost:8181",
	"basePath": "/",
	"securityDefinitions": {},
	"schemes": [
	  "http"
	],
	"consumes": [
	  "application/json"
	],
	"produces": [
	  "application/json"
	],
	"paths": {
	  "/customer": {
		"get": {
		  "summary": "List Customer",
		  "tags": [
			"Costumer"
		  ],
		  "operationId": "ListCustomer",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "Authorization",
			  "in": "header",
			  "required": false,
			  "default": "Bearer {token}",
			  "type": "string"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  }
		},
		"post": {
		  "summary": "New Customer",
		  "tags": [
			"Costumer"
		  ],
		  "operationId": "NewCustomer",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "Authorization",
			  "in": "header",
			  "required": false,
			  "default": "Bearer {token}",
			  "type": "string"
			},
			{
			  "name": "Body",
			  "in": "body",
			  "required": true,
			  "description": "",
			  "schema": {
				"$ref": "#/definitions/NewCustomerRequest"
			  }
			}
		  ],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  }
		},
		"put": {
		  "summary": "Update Customer",
		  "tags": [
			"Costumer"
		  ],
		  "operationId": "UpdateCustomer",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "Authorization",
			  "in": "header",
			  "required": false,
			  "default": "Bearer {token}",
			  "type": "string"
			},
			{
			  "name": "Body",
			  "in": "body",
			  "required": true,
			  "description": "",
			  "schema": {
				"example": "{\r\n   \"id\":5\",\r\n   \"name\":\"Doni\",\r\n   \"email\": \"email@gmail.com\",\r\n   \"address\":\"Palembang\"\r\n}",
				"type": "string"
			  }
			}
		  ],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  }
		}
	  },
	  "/customer/{id}": {
		"get": {
		  "summary": "Find Customer By ID",
		  "tags": [
			"Costumer"
		  ],
		  "operationId": "FindCustomerByID",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "Authorization",
			  "in": "header",
			  "required": false,
			  "default": "Bearer {token}",
			  "type": "string"
			},
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "type": "string",
			  "description": ""
			}
		  ],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  }
		},
		"delete": {
		  "summary": "Delete Customer",
		  "tags": [
			"Costumer"
		  ],
		  "operationId": "DeleteCustomer",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "Authorization",
			  "in": "header",
			  "required": false,
			  "default": "Bearer {token}",
			  "type": "string"
			},
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "type": "string",
			  "description": ""
			}
		  ],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  }
		}
	  },
	  "/enigma/auth": {
		"post": {
		  "summary": "Auth",
		  "tags": [
			"Authenticate"
		  ],
		  "operationId": "Auth",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "Authorization",
			  "in": "header",
			  "required": false,
			  "default": "Bearer {token}",
			  "type": "string"
			},
			{
			  "name": "Body",
			  "in": "body",
			  "required": true,
			  "description": "",
			  "schema": {
				"$ref": "#/definitions/AuthRequest"
			  }
			}
		  ],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  }
		}
	  },
	  "/enigma/protected/user": {
		"get": {
		  "summary": "Test Login User",
		  "tags": [
			"Authenticate"
		  ],
		  "operationId": "TestLoginUser",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [],
		  "responses": {
			"200": {
			  "description": "",
			  "headers": {}
			}
		  },
		  "security": []
		}
	  }
	},
	"definitions": {
	  "NewCustomerRequest": {
		"title": "NewCustomerRequest",
		"example": {
		  "id": "1",
		  "name": "Vicky",
		  "email": "vick@gmail.com",
		  "address": "jakarta"
		},
		"type": "object",
		"properties": {
		  "id": {
			"type": "string"
		  },
		  "name": {
			"type": "string"
		  },
		  "email": {
			"type": "string"
		  },
		  "address": {
			"type": "string"
		  }
		},
		"required": [
		  "id",
		  "name",
		  "email",
		  "address"
		]
	  },
	  "AuthRequest": {
		"title": "AuthRequest",
		"example": {
		  "username": "enigma",
		  "password": "123"
		},
		"type": "object",
		"properties": {
		  "username": {
			"type": "string"
		  },
		  "password": {
			"type": "string"
		  }
		},
		"required": [
		  "username",
		  "password"
		]
	  }
	},
	"security": [],
	"tags": [
	  {
		"name": "Costumer"
	  },
	  {
		"name": "Authenticate"
	  }
	]
  }`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Warung Makan API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
