// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-10-07 12:38:52.0428924 +0100 BST m=+0.368133701

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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/nymtech/nym-directory/license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/healthcheck": {
            "get": {
                "description": "Returns a 200 if the directory server is available. Good route to use for automated monitoring.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Lets the directory server tell the world it's alive.",
                "operationId": "healthCheck",
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/metrics/mixes": {
            "get": {
                "description": "For demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Lists mixnode activity in the past 3 seconds",
                "operationId": "listMixMetrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MixMetric"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "For demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Create a metric detailing how many messages a given mixnode sent and received",
                "operationId": "createMixMetric",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixMetric"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining": {
            "post": {
                "description": "Nym network monitor sends packets through the system and checks if they make it. The network monitor then hits this method to report whether the node was up at a given time.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Lets the network monitor create a new uptime status for a mix",
                "operationId": "addMixStatus",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixStatus"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/batch": {
            "post": {
                "description": "Nym network monitor sends packets through the system and checks if they make it. The network monitor then hits this method to report whether nodes were up at a given time.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Lets the network monitor create a new uptime status for multiple mixes",
                "operationId": "addBatchMixStatus",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BatchMixStatus"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/{pubkey}/history": {
            "get": {
                "description": "Lists all mixnode statuses for a given node pubkey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Lists mixnode activity",
                "operationId": "listMixStatuses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mixnode Pubkey",
                        "name": "pubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MixStatus"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/{pubkey}/report": {
            "get": {
                "description": "Provides summary uptime statistics for last 5 minutes, day, week, and month",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Retrieves a summary report of historical mix status",
                "operationId": "getMixStatusReport",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mixnode Pubkey",
                        "name": "pubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/allow": {
            "post": {
                "description": "Sometimes when a node isn't working we need to temporarily remove it. This allows us to re-enable it once it's working again.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Removes a disallowed node from the disallowed nodes list",
                "operationId": "allow",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixNodeID"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/coconodes": {
            "post": {
                "description": "Nym Coconut nodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets a coconut node tell the directory server it's alive",
                "operationId": "addCocoNode",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CocoHostInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/disallow": {
            "post": {
                "description": "Sometimes when a node isn't working we need to temporarily remove it from use so that it doesn't mess up QoS for the whole network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Takes a node out of the regular topology and puts it in the disallowed nodes list",
                "operationId": "disallow",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixNodeID"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/disallowed": {
            "get": {
                "description": "Sometimes we need to take mixnodes out of the network for repair. This shows which ones are currently disallowed.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lists Nym mixnodes that are currently disallowed",
                "operationId": "disallowed",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MixNodePresence"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/gateways": {
            "post": {
                "description": "Nym mix gateways can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets a gateway tell the directory server it's alive",
                "operationId": "addGateway",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GatewayHostInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/mixnodes": {
            "post": {
                "description": "Nym mixnodes can ping this method to let the directory server know they're up. We can then use this info to create topologies of the overall Nym network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets mixnode a node tell the directory server it's alive",
                "operationId": "addMixNode",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixHostInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/topology": {
            "get": {
                "description": "Nym nodes periodically ping the directory server to register that they're alive. This method provides a list of nodes which have been most recently seen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lists which Nym mixnodes, providers, gateways, and coconodes are alive",
                "operationId": "topology",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Topology"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BatchMixStatus": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MixStatus"
                    }
                }
            }
        },
        "models.CocoHostInfo": {
            "type": "object",
            "required": [
                "pubKey",
                "type",
                "version"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.CocoPresence": {
            "type": "object",
            "required": [
                "lastSeen",
                "pubKey",
                "type",
                "version"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.GatewayHostInfo": {
            "type": "object",
            "required": [
                "clientListener",
                "identityKey",
                "mixnetListener",
                "sphinxKey",
                "version"
            ],
            "properties": {
                "clientListener": {
                    "type": "string"
                },
                "identityKey": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "mixnetListener": {
                    "type": "string"
                },
                "sphinxKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.GatewayPresence": {
            "type": "object",
            "required": [
                "clientListener",
                "identityKey",
                "lastSeen",
                "mixnetListener",
                "sphinxKey",
                "version"
            ],
            "properties": {
                "clientListener": {
                    "type": "string"
                },
                "identityKey": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "mixnetListener": {
                    "type": "string"
                },
                "sphinxKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.MixHostInfo": {
            "type": "object",
            "required": [
                "layer",
                "pubKey",
                "version"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "layer": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.MixMetric": {
            "type": "object",
            "required": [
                "pubKey",
                "received",
                "sent"
            ],
            "properties": {
                "pubKey": {
                    "type": "string"
                },
                "received": {
                    "type": "integer"
                },
                "sent": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.MixNodeID": {
            "type": "object",
            "properties": {
                "pubKey": {
                    "type": "string"
                }
            }
        },
        "models.MixNodePresence": {
            "type": "object",
            "required": [
                "lastSeen",
                "layer",
                "pubKey",
                "version"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "lastSeen": {
                    "description": "MixStatusReport",
                    "type": "integer"
                },
                "layer": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.MixProviderPresence": {
            "type": "object",
            "required": [
                "lastSeen",
                "pubKey",
                "version"
            ],
            "properties": {
                "clientListener": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "mixnetListener": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.MixStatus": {
            "type": "object",
            "required": [
                "ipVersion",
                "pubKey",
                "up"
            ],
            "properties": {
                "ipVersion": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "up": {
                    "type": "boolean"
                }
            }
        },
        "models.Topology": {
            "type": "object",
            "properties": {
                "cocoNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CocoPresence"
                    }
                },
                "gatewayNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.GatewayPresence"
                    }
                },
                "mixNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MixNodePresence"
                    }
                },
                "mixProviderNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MixProviderPresence"
                    }
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
	Version:     "0.9.0-dev",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Nym Directory API",
	Description: "This is a temporarily centralized directory/PKI/metrics API to allow us to get the other Nym node types running. Its functionality will eventually be folded into other parts of Nym.",
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
