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
	"basePath": "/v1",
	"securityDefinitions": {},
	"schemes": [
	  "https"
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
		}
	  },
	  "/private/customer": {
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
		},
		"get": {
		  "summary": "Test (read only)",
		  "tags": [
			"Authenticate"
		  ],
		  "operationId": "Test(readonly)",
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
	  },
	  "/private/customer/{id}": {
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
	  "/auth": {
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
		  },
		  "security": []
		}
	  },
	  "/private/order": {
		"get": {
		  "summary": "List Menu",
		  "tags": [
			"Menu"
		  ],
		  "operationId": "ListMenu",
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
		  "summary": "New Order",
		  "tags": [
			"order"
		  ],
		  "operationId": "NewOrder",
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
				"$ref": "#/definitions/NewOrderRequest"
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
	  "/private/menu": {
		"post": {
		  "summary": "New menu",
		  "tags": [
			"Menu"
		  ],
		  "operationId": "Newmenu",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "consumes": [
			"application/x-www-form-urlencoded"
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
			  "name": "menu_name",
			  "in": "formData",
			  "required": true,
			  "type": "string",
			  "description": ""
			},
			{
			  "name": "price",
			  "in": "formData",
			  "required": true,
			  "type": "integer",
			  "format": "int32",
			  "description": ""
			},
			{
			  "name": "images",
			  "in": "formData",
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
		"put": {
		  "summary": "Update Menu",
		  "tags": [
			"Menu"
		  ],
		  "operationId": "UpdateMenu",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "consumes": [
			"application/x-www-form-urlencoded"
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
			  "in": "formData",
			  "required": true,
			  "type": "integer",
			  "format": "int32",
			  "description": ""
			},
			{
			  "name": "menu_name",
			  "in": "formData",
			  "required": true,
			  "type": "string",
			  "description": ""
			},
			{
			  "name": "price",
			  "in": "formData",
			  "required": true,
			  "type": "integer",
			  "format": "int32",
			  "description": ""
			},
			{
			  "name": "images",
			  "in": "formData",
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
		"get": {
		  "summary": "List Transaction",
		  "tags": [
			"order"
		  ],
		  "operationId": "ListTransaction",
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
	  },
	  "/private/menu/{id}": {
		"get": {
		  "summary": "Find Menu By ID",
		  "tags": [
			"Menu"
		  ],
		  "operationId": "FindMenuByID",
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
	  "/private/menu/11": {
		"delete": {
		  "summary": "Delete Menu",
		  "tags": [
			"Menu"
		  ],
		  "operationId": "DeleteMenu",
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
		}
	  },
	  "/private/menu/images/{nama_file}": {
		"get": {
		  "summary": "Download Menu Image",
		  "tags": [
			"Menu"
		  ],
		  "operationId": "DownloadMenuImage",
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
			  "name": "nama_file",
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
	  "/private/order/{id}": {
		"get": {
		  "summary": "Find Transaction By ID",
		  "tags": [
			"order"
		  ],
		  "operationId": "FindTransactionByID",
		  "deprecated": false,
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
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
		  },
		  "security": []
		},
		"delete": {
		  "summary": "Delete Transaction",
		  "tags": [
			"order"
		  ],
		  "operationId": "DeleteTransaction",
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
	  },
	  "NewOrderRequest": {
		"title": "NewOrderRequest",
		"example": {
		  "id": "1234567892",
		  "customer_id": "1",
		  "table_id": 1,
		  "paid_status": true,
		  "total_price": 50000,
		  "order_detail_id": 56,
		  "details": {
			"order_details": [
			  {
				"id": 1234567894,
				"menu_id": 1,
				"quantity": 2,
				"order_id": 56
			  },
			  {
				"id": 1234567895,
				"menu_id": 2,
				"quantity": 1,
				"order_id": 56
			  }
			]
		  }
		},
		"type": "object",
		"properties": {
		  "id": {
			"type": "string"
		  },
		  "customer_id": {
			"type": "string"
		  },
		  "table_id": {
			"type": "integer",
			"format": "int32"
		  },
		  "paid_status": {
			"type": "boolean"
		  },
		  "total_price": {
			"type": "integer",
			"format": "int32"
		  },
		  "order_detail_id": {
			"type": "integer",
			"format": "int32"
		  },
		  "details": {
			"$ref": "#/definitions/Details"
		  }
		},
		"required": [
		  "id",
		  "customer_id",
		  "table_id",
		  "paid_status",
		  "total_price",
		  "order_detail_id",
		  "details"
		]
	  },
	  "Details": {
		"title": "Details",
		"example": {
		  "order_details": [
			{
			  "id": 1234567894,
			  "menu_id": 1,
			  "quantity": 2,
			  "order_id": 56
			},
			{
			  "id": 1234567895,
			  "menu_id": 2,
			  "quantity": 1,
			  "order_id": 56
			}
		  ]
		},
		"type": "object",
		"properties": {
		  "order_details": {
			"type": "array",
			"items": {
			  "$ref": "#/definitions/OrderDetail"
			}
		  }
		},
		"required": [
		  "order_details"
		]
	  },
	  "OrderDetail": {
		"title": "OrderDetail",
		"example": {
		  "id": 1234567894,
		  "menu_id": 1,
		  "quantity": 2,
		  "order_id": 56
		},
		"type": "object",
		"properties": {
		  "id": {
			"type": "integer",
			"format": "int32"
		  },
		  "menu_id": {
			"type": "integer",
			"format": "int32"
		  },
		  "quantity": {
			"type": "integer",
			"format": "int32"
		  },
		  "order_id": {
			"type": "integer",
			"format": "int32"
		  }
		},
		"required": [
		  "id",
		  "menu_id",
		  "quantity",
		  "order_id"
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
	  },
	  {
		"name": "Menu"
	  },
	  {
		"name": "order"
	  }
	]
  }`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8181",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Warung Makan API",
	Description:      "This is a simple warung makan server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
