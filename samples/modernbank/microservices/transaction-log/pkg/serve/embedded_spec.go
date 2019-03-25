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
    "description": "This is the Transaction Log Microservice, an immutable append-only Transaction Log in Modern Bank.",
    "title": "Transaction Log",
    "version": "1.0.0"
  },
  "host": "transaction-log",
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
    "/transactions": {
      "post": {
        "description": "Sends money from one account to another",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Sends money from one account to another",
        "operationId": "CreateTransaction",
        "parameters": [
          {
            "description": "Transaction to create",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newtransaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/transaction"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "400": {
            "description": "Nice try! You can't send negative amounts...",
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
    "/transactions/account/{receiver}/received": {
      "get": {
        "description": "Lists all transactions sent to a given account",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Lists all transactions sent to a given account",
        "operationId": "listTransactionsReceived",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Account number that received the transactions",
            "name": "receiver",
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
                "$ref": "#/definitions/transaction"
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
            "description": "No transactions found",
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
    "/transactions/account/{sender}/sent": {
      "get": {
        "description": "Lists all transactions sent from a given account",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Lists all transactions sent from a given account",
        "operationId": "listTransactionsSent",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Account number that made the transactions",
            "name": "sender",
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
                "$ref": "#/definitions/transaction"
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
            "description": "No transactions found",
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
    "newtransaction": {
      "type": "object",
      "required": [
        "sender",
        "receiver",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "number",
          "format": "currency"
        },
        "receiver": {
          "type": "integer",
          "format": "int64"
        },
        "sender": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "transaction": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "sender",
            "receiver",
            "amount"
          ],
          "properties": {
            "amount": {
              "type": "number",
              "format": "currency"
            },
            "receiver": {
              "type": "integer",
              "format": "int64"
            },
            "sender": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        {
          "type": "object",
          "required": [
            "id"
          ],
          "properties": {
            "id": {
              "type": "string"
            }
          }
        }
      ]
    }
  },
  "tags": [
    {
      "description": "Operations about transactions",
      "name": "transaction-log"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Transaction Log Microservice, an immutable append-only Transaction Log in Modern Bank.",
    "title": "Transaction Log",
    "version": "1.0.0"
  },
  "host": "transaction-log",
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
    "/transactions": {
      "post": {
        "description": "Sends money from one account to another",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Sends money from one account to another",
        "operationId": "CreateTransaction",
        "parameters": [
          {
            "description": "Transaction to create",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newtransaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/transaction"
            },
            "headers": {
              "version": {
                "type": "string",
                "description": "Version of the microservice that responded"
              }
            }
          },
          "400": {
            "description": "Nice try! You can't send negative amounts...",
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
    "/transactions/account/{receiver}/received": {
      "get": {
        "description": "Lists all transactions sent to a given account",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Lists all transactions sent to a given account",
        "operationId": "listTransactionsReceived",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Account number that received the transactions",
            "name": "receiver",
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
                "$ref": "#/definitions/transaction"
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
            "description": "No transactions found",
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
    "/transactions/account/{sender}/sent": {
      "get": {
        "description": "Lists all transactions sent from a given account",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Lists all transactions sent from a given account",
        "operationId": "listTransactionsSent",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Account number that made the transactions",
            "name": "sender",
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
                "$ref": "#/definitions/transaction"
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
            "description": "No transactions found",
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
    "newtransaction": {
      "type": "object",
      "required": [
        "sender",
        "receiver",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "number",
          "format": "currency"
        },
        "receiver": {
          "type": "integer",
          "format": "int64"
        },
        "sender": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "transaction": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "sender",
            "receiver",
            "amount"
          ],
          "properties": {
            "amount": {
              "type": "number",
              "format": "currency"
            },
            "receiver": {
              "type": "integer",
              "format": "int64"
            },
            "sender": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        {
          "type": "object",
          "required": [
            "id"
          ],
          "properties": {
            "id": {
              "type": "string"
            }
          }
        }
      ]
    }
  },
  "tags": [
    {
      "description": "Operations about transactions",
      "name": "transaction-log"
    }
  ]
}`))
}
