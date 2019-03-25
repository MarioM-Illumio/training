// Code generated by go-swagger; DO NOT EDIT.

package serve

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the User Microservice, responsible for managing Users in Modern Bank.",
    "title": "User",
    "version": "1.0.0"
  },
  "host": "users",
  "basePath": "/v1",
  "paths": {
    "/health": {
      "post": {
        "description": "returns 200",
        "tags": [
          "health"
        ],
        "summary": "returns 200 to prove the service is alive",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "description": "Creates a new user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Create a new user identity for a customer",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "409": {
            "description": "User alreadys exists",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByUserName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "User not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    },
    "/users/{username}/accounts": {
      "get": {
        "description": "Lists all accounts for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Lists all accounts for a given customer",
        "operationId": "listAccounts",
        "parameters": [
          {
            "type": "string",
            "description": "Username to fetch the accounts of",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/account"
              }
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "Owner not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "account": {
      "type": "object",
      "required": [
        "number",
        "balance",
        "owner",
        "type"
      ],
      "properties": {
        "balance": {
          "type": "number",
          "format": "currency"
        },
        "number": {
          "type": "integer",
          "format": "int64"
        },
        "owner": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "cash",
            "saving"
          ]
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "username",
        "firstName",
        "lastName",
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about users",
      "name": "user"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the User Microservice, responsible for managing Users in Modern Bank.",
    "title": "User",
    "version": "1.0.0"
  },
  "host": "users",
  "basePath": "/v1",
  "paths": {
    "/health": {
      "post": {
        "description": "returns 200",
        "tags": [
          "health"
        ],
        "summary": "returns 200 to prove the service is alive",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    },
    "/users": {
      "post": {
        "description": "Creates a new user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Create a new user identity for a customer",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "409": {
            "description": "User alreadys exists",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByUserName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "User not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    },
    "/users/{username}/accounts": {
      "get": {
        "description": "Lists all accounts for a given customer",
        "produces": [
          "application/json"
        ],
        "tags": [
          "accounts"
        ],
        "summary": "Lists all accounts for a given customer",
        "operationId": "listAccounts",
        "parameters": [
          {
            "type": "string",
            "description": "Username to fetch the accounts of",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/account"
              }
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "404": {
            "description": "Owner not found",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "500": {
            "description": "Internal server error",
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "account": {
      "type": "object",
      "required": [
        "number",
        "balance",
        "owner",
        "type"
      ],
      "properties": {
        "balance": {
          "type": "number",
          "format": "currency"
        },
        "number": {
          "type": "integer",
          "format": "int64"
        },
        "owner": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "cash",
            "saving"
          ]
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "username",
        "firstName",
        "lastName",
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about users",
      "name": "user"
    }
  ]
}`))
}
