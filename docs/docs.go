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
        "/calculate/fee": {
            "post": {
                "description": "Calculate the fee based on the provided amount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculate"
                ],
                "summary": "Calculate fee",
                "parameters": [
                    {
                        "description": "Calculate Fee Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.CalculateFeeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Calculation result",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/calculate/help": {
            "get": {
                "description": "Calculate help information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculate"
                ],
                "summary": "Calculate help information",
                "responses": {
                    "200": {
                        "description": "Successful calculation result",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/calculate/reward": {
            "post": {
                "description": "Calculate the reward based on the provided stake and number of days",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculate"
                ],
                "summary": "Calculate reward",
                "parameters": [
                    {
                        "description": "Calculate Reward Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.CalculateRewardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Calculation result",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/commands": {
            "get": {
                "description": "Retrieve a list of all available commands",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Commands"
                ],
                "summary": "Get available commands",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.GetCommandsResponse"
                        }
                    }
                }
            }
        },
        "/help": {
            "get": {
                "description": "This endpoint runs the 'help' command and returns the result, including color, title, error message, and success status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Help"
                ],
                "summary": "Executes a help command and returns the result.",
                "responses": {
                    "200": {
                        "description": "Help command result",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Result": {
                                            "$ref": "#/definitions/http.CommandResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/market/help": {
            "get": {
                "description": "This endpoint processes the \"market help\" command and returns relevant details about the market command.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Market"
                ],
                "summary": "Get help for the market command",
                "responses": {
                    "200": {
                        "description": "Command help details",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Result": {
                                            "$ref": "#/definitions/http.CommandResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/market/price": {
            "get": {
                "description": "Retrieves the current market price information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Market"
                ],
                "summary": "Get Market Price",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the market price",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/http.CommandResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/network/health": {
            "get": {
                "description": "Retrieve the health status of the network",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Network"
                ],
                "summary": "Get network health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/network/help": {
            "get": {
                "description": "Provides network help information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Network"
                ],
                "summary": "Provides network help information",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved network help",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/http.CommandResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/network/node-info": {
            "post": {
                "description": "Retrieves information about a network node based on the provided validator address.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Network"
                ],
                "summary": "Get Network Node Information",
                "parameters": [
                    {
                        "description": "Validator Address",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.NetworkNodeInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The result of the command",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/network/status": {
            "post": {
                "description": "Retrieves the current status of the network node, including command results such as color, title, error message, and success status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Network"
                ],
                "summary": "Get network node status",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved network node status",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/phoenix/faucet": {
            "post": {
                "description": "Initiates a faucet request for the Phoenix",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phoenix"
                ],
                "summary": "Initiates a faucet request for the Phoenix",
                "parameters": [
                    {
                        "description": "Faucet request with address",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.PhoenixFaucetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response containing faucet status",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/phoenix/help": {
            "get": {
                "description": "Executes Phoenix help command.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phoenix"
                ],
                "summary": "Executes Phoenix help command",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.BasicResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/http.CommandResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/phoenix/status": {
            "get": {
                "description": "Get Phoenix Status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Phoenix"
                ],
                "summary": "Get Phoenix Status",
                "responses": {
                    "200": {
                        "description": "Status of the Phoenix application",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/voucher/claim": {
            "post": {
                "description": "This endpoint allows users to claim a voucher by providing a valid code and their address.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voucher"
                ],
                "summary": "Claim a voucher using the provided code and address.",
                "parameters": [
                    {
                        "description": "Voucher Claim Request",
                        "name": "voucherClaimRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.VoucherClaimRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Voucher claim successful",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/voucher/create-one": {
            "post": {
                "description": "Creates a voucher.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voucher"
                ],
                "summary": "Creates a new voucher",
                "parameters": [
                    {
                        "description": "Voucher details",
                        "name": "voucherCreateOneRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.VoucherCreateOneRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Voucher creation successful",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/voucher/help": {
            "get": {
                "description": "Executes the 'voucher help' command and returns the result including color, title, error, message, and success status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voucher"
                ],
                "summary": "Get help information for the voucher command",
                "responses": {
                    "200": {
                        "description": "Successful response with command result",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/voucher/status": {
            "post": {
                "description": "Accepts a voucher code and returns the status of the voucher, including title, message, error, color, and success information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voucher"
                ],
                "summary": "Check the status of a voucher",
                "parameters": [
                    {
                        "description": "Voucher Status Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.VoucherStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Voucher status response",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        },
        "/vouchers/create-bulk": {
            "post": {
                "description": "This API allows you to create multiple vouchers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voucher"
                ],
                "summary": "Create vouchers in bulk",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File containing voucher data",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Notification flag (optional)",
                        "name": "notify",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Command result containing the status of the bulk voucher creation",
                        "schema": {
                            "$ref": "#/definitions/http.BasicResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "command.InputBox": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "InputBoxText",
                "InputBoxNumber",
                "InputBoxFile",
                "InputBoxAmount",
                "InputBoxToggle"
            ]
        },
        "http.BasicResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/http.CommandResult"
                }
            }
        },
        "http.CalculateFeeRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                }
            }
        },
        "http.CalculateRewardRequest": {
            "type": "object",
            "properties": {
                "days": {
                    "type": "string"
                },
                "stake": {
                    "type": "string"
                }
            }
        },
        "http.CommandArgsResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "optional": {
                    "type": "boolean"
                },
                "type": {
                    "$ref": "#/definitions/command.InputBox"
                }
            }
        },
        "http.CommandResponse": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.CommandArgsResponse"
                    }
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "subCommands": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.CommandResponse"
                    }
                }
            }
        },
        "http.CommandResult": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "successful": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "http.GetCommandsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.CommandResponse"
                    }
                }
            }
        },
        "http.NetworkNodeInfoRequest": {
            "type": "object",
            "properties": {
                "validatorAddress": {
                    "type": "string"
                }
            }
        },
        "http.PhoenixFaucetRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                }
            }
        },
        "http.VoucherClaimRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                }
            }
        },
        "http.VoucherCreateOneRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "recipient": {
                    "type": "string"
                },
                "validMonths": {
                    "type": "string"
                }
            }
        },
        "http.VoucherStatusRequest": {
            "type": "object",
            "properties": {
                "code": {
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
	BasePath:         "/api/v1",
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
