// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package rest

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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "libopenstorage autopilot API",
    "title": "autopilot",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localhost:9000",
  "basePath": "/api/v1",
  "paths": {
    "/collectors": {
      "get": {
        "description": "Returns an array of telemetry collectors defined in the system",
        "summary": "Get a list of telemetry collectors",
        "operationId": "collectorList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Collector"
              }
            }
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      }
    },
    "/collectors/{collector}/poll": {
      "get": {
        "description": "Poll a collector for the given period directly",
        "summary": "Poll a collector",
        "operationId": "collectorPoll",
        "responses": {
          "201": {
            "description": "No Content"
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the collector name",
          "name": "collector",
          "in": "path",
          "required": true
        }
      ]
    },
    "/emitters": {
      "get": {
        "description": "Returns an array of telemetry emitters defined in the system",
        "summary": "Get a list of telemetry emitters",
        "operationId": "emitterList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Emitter"
              }
            }
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      }
    },
    "/providers": {
      "get": {
        "description": "Returns an array of telemetry providers defined in the system",
        "summary": "Get a list of telemetry providers",
        "operationId": "providerList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Provider"
              }
            }
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      }
    },
    "/providers/{provider}/query": {
      "get": {
        "description": "Query a provider directly",
        "summary": "Query a collector",
        "operationId": "providerQuery",
        "parameters": [
          {
            "type": "string",
            "format": "date-time",
            "name": "start_date",
            "in": "query"
          },
          {
            "type": "string",
            "format": "date-time",
            "name": "end_date",
            "in": "query"
          },
          {
            "type": "string",
            "name": "query",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Collector"
              }
            }
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the provider name",
          "name": "provider",
          "in": "path",
          "required": true
        }
      ]
    },
    "/providers/{provider}/recommend": {
      "get": {
        "description": "Create a new telemetry sample from the provided definition and get recommendations",
        "summary": "Post a telemetry sample and get recommendations",
        "operationId": "recommendationsGet",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "csv",
            "description": "The pre-defined rules to apply",
            "name": "rules",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Recommendation"
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the provider name",
          "name": "provider",
          "in": "path",
          "required": true
        }
      ]
    },
    "/rules": {
      "get": {
        "description": "Returns an array of telemetry rules defined in the system",
        "summary": "Get a list of telemetry rules",
        "operationId": "ruleList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Rule"
              }
            }
          },
          "500": {
            "$ref": "#/responses/ServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "Collector": {
      "description": "A collector pulls data from a provider at regular intervals and stores the data in an autopilot format for the ML engine.\n",
      "properties": {
        "interval": {
          "description": "The interval for the collecto to run and collect",
          "type": "string",
          "default": "24h"
        },
        "name": {
          "description": "The collector name",
          "type": "string"
        },
        "params": {
          "description": "The collector provider params",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "provider": {
          "description": "The provider name",
          "type": "string"
        }
      }
    },
    "Emitter": {
      "description": "An emitter emits recommendations to a target\n",
      "properties": {
        "name": {
          "description": "The emitter name",
          "type": "string"
        },
        "params": {
          "description": "json data object",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "type": {
          "description": "The emitter type",
          "type": "string",
          "enum": [
            "mqtt"
          ]
        }
      }
    },
    "Error": {
      "description": "Common Error Model",
      "type": "object",
      "properties": {
        "code": {
          "description": "The error code",
          "type": "string",
          "x-nullable": true
        },
        "detail": {
          "description": "The error details",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-nullable": true,
          "x-omitempty": true
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "Monitor": {
      "description": "A monitor executes and analyzes rules at the specified interval, emmiting alerts\n",
      "properties": {
        "interval": {
          "description": "The interval to monitor",
          "type": "string",
          "default": "15m"
        },
        "name": {
          "description": "The monitor name",
          "type": "string"
        },
        "params": {
          "description": "the monitor provider additional params",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "provider": {
          "description": "The provider to monitor",
          "type": "string"
        },
        "rules": {
          "description": "The rules to execute on the provider",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Prometheus": {
      "description": "Prometheus provider configuration",
      "allOf": [
        {
          "$ref": "#/definitions/Provider"
        },
        {
          "properties": {
            "url": {
              "description": "The prometheus host url",
              "type": "string"
            }
          }
        }
      ]
    },
    "Proposal": {
      "description": "A proposal is a formatted propsal object\n",
      "properties": {
        "action": {
          "description": "The proposed action to take to resolve the issue",
          "type": "string"
        },
        "cluster": {
          "description": "The cluster id",
          "type": "string"
        },
        "issue": {
          "description": "Issue from the rule that describes the reason for this proposal",
          "type": "string"
        },
        "node": {
          "description": "The node id",
          "type": "string"
        },
        "rule": {
          "description": "The rule that triggered the proposal",
          "type": "string"
        },
        "volume": {
          "description": "The volume id",
          "type": "string"
        }
      }
    },
    "Provider": {
      "description": "A provider defines a provider configuration\n",
      "properties": {
        "name": {
          "description": "The provider instance name",
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/ProviderType"
        }
      },
      "discriminator": "type"
    },
    "ProviderType": {
      "description": "Provider types",
      "type": "string",
      "enum": [
        "Prometheus"
      ]
    },
    "Recommendation": {
      "description": "A recommendation is a list of recommended arbitrations to be emitted by the system\n",
      "properties": {
        "proposals": {
          "description": "The recommendation values mapping rule.name -\u003e formatted proposal",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Proposal"
          }
        },
        "timestamp": {
          "description": "The recommendation timestamp",
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "Rule": {
      "description": "An rule is a contraint expression that checked in the system against\n",
      "properties": {
        "expr": {
          "description": "The expression to match or query to make",
          "type": "string",
          "example": "100 * (px_volume_usage_bytes / px_volume_capacity_bytes) \u003e 80"
        },
        "for": {
          "description": "The duration/interval the expression must be valid for in seconds",
          "type": "integer",
          "format": "int64",
          "example": 3600
        },
        "issue": {
          "description": "The issue template",
          "type": "string",
          "example": "Portworx volume {{$labels.volume}} usage on {{$labels.host}} is high."
        },
        "name": {
          "description": "the rule description",
          "type": "string"
        },
        "proposal": {
          "description": "The proposal template",
          "type": "string",
          "example": "Add additional storage node to {{$labels.cluster}}"
        },
        "severity": {
          "type": "string",
          "enum": [
            "warning",
            "error",
            "critical"
          ]
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "BadRequest",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "invalid parameter"
        }
      }
    },
    "NotFound": {
      "description": "NotFound",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "object not found"
        }
      }
    },
    "ServerError": {
      "description": "ServerError",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "internal server error"
        }
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "access denied"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "security": [
    {
      "basicAuth": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "libopenstorage autopilot API",
    "title": "autopilot",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localhost:9000",
  "basePath": "/api/v1",
  "paths": {
    "/collectors": {
      "get": {
        "description": "Returns an array of telemetry collectors defined in the system",
        "summary": "Get a list of telemetry collectors",
        "operationId": "collectorList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Collector"
              }
            }
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      }
    },
    "/collectors/{collector}/poll": {
      "get": {
        "description": "Poll a collector for the given period directly",
        "summary": "Poll a collector",
        "operationId": "collectorPoll",
        "responses": {
          "201": {
            "description": "No Content"
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the collector name",
          "name": "collector",
          "in": "path",
          "required": true
        }
      ]
    },
    "/emitters": {
      "get": {
        "description": "Returns an array of telemetry emitters defined in the system",
        "summary": "Get a list of telemetry emitters",
        "operationId": "emitterList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Emitter"
              }
            }
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      }
    },
    "/providers": {
      "get": {
        "description": "Returns an array of telemetry providers defined in the system",
        "summary": "Get a list of telemetry providers",
        "operationId": "providerList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Provider"
              }
            }
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      }
    },
    "/providers/{provider}/query": {
      "get": {
        "description": "Query a provider directly",
        "summary": "Query a collector",
        "operationId": "providerQuery",
        "parameters": [
          {
            "type": "string",
            "format": "date-time",
            "name": "start_date",
            "in": "query"
          },
          {
            "type": "string",
            "format": "date-time",
            "name": "end_date",
            "in": "query"
          },
          {
            "type": "string",
            "name": "query",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Collector"
              }
            }
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the provider name",
          "name": "provider",
          "in": "path",
          "required": true
        }
      ]
    },
    "/providers/{provider}/recommend": {
      "get": {
        "description": "Create a new telemetry sample from the provided definition and get recommendations",
        "summary": "Post a telemetry sample and get recommendations",
        "operationId": "recommendationsGet",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "csv",
            "description": "The pre-defined rules to apply",
            "name": "rules",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Recommendation"
              }
            }
          },
          "400": {
            "description": "BadRequest",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "invalid parameter"
              }
            }
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the provider name",
          "name": "provider",
          "in": "path",
          "required": true
        }
      ]
    },
    "/rules": {
      "get": {
        "description": "Returns an array of telemetry rules defined in the system",
        "summary": "Get a list of telemetry rules",
        "operationId": "ruleList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Rule"
              }
            }
          },
          "500": {
            "description": "ServerError",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "examples": {
              "application/json": {
                "message": "internal server error"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Collector": {
      "description": "A collector pulls data from a provider at regular intervals and stores the data in an autopilot format for the ML engine.\n",
      "properties": {
        "interval": {
          "description": "The interval for the collecto to run and collect",
          "type": "string",
          "default": "24h"
        },
        "name": {
          "description": "The collector name",
          "type": "string"
        },
        "params": {
          "description": "The collector provider params",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "provider": {
          "description": "The provider name",
          "type": "string"
        }
      }
    },
    "Emitter": {
      "description": "An emitter emits recommendations to a target\n",
      "properties": {
        "name": {
          "description": "The emitter name",
          "type": "string"
        },
        "params": {
          "description": "json data object",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "type": {
          "description": "The emitter type",
          "type": "string",
          "enum": [
            "mqtt"
          ]
        }
      }
    },
    "Error": {
      "description": "Common Error Model",
      "type": "object",
      "properties": {
        "code": {
          "description": "The error code",
          "type": "string",
          "x-nullable": true
        },
        "detail": {
          "description": "The error details",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-nullable": true,
          "x-omitempty": true
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "Monitor": {
      "description": "A monitor executes and analyzes rules at the specified interval, emmiting alerts\n",
      "properties": {
        "interval": {
          "description": "The interval to monitor",
          "type": "string",
          "default": "15m"
        },
        "name": {
          "description": "The monitor name",
          "type": "string"
        },
        "params": {
          "description": "the monitor provider additional params",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "provider": {
          "description": "The provider to monitor",
          "type": "string"
        },
        "rules": {
          "description": "The rules to execute on the provider",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Prometheus": {
      "description": "Prometheus provider configuration",
      "allOf": [
        {
          "$ref": "#/definitions/Provider"
        },
        {
          "properties": {
            "url": {
              "description": "The prometheus host url",
              "type": "string"
            }
          }
        }
      ]
    },
    "Proposal": {
      "description": "A proposal is a formatted propsal object\n",
      "properties": {
        "action": {
          "description": "The proposed action to take to resolve the issue",
          "type": "string"
        },
        "cluster": {
          "description": "The cluster id",
          "type": "string"
        },
        "issue": {
          "description": "Issue from the rule that describes the reason for this proposal",
          "type": "string"
        },
        "node": {
          "description": "The node id",
          "type": "string"
        },
        "rule": {
          "description": "The rule that triggered the proposal",
          "type": "string"
        },
        "volume": {
          "description": "The volume id",
          "type": "string"
        }
      }
    },
    "Provider": {
      "description": "A provider defines a provider configuration\n",
      "properties": {
        "name": {
          "description": "The provider instance name",
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/ProviderType"
        }
      },
      "discriminator": "type"
    },
    "ProviderType": {
      "description": "Provider types",
      "type": "string",
      "enum": [
        "Prometheus"
      ]
    },
    "Recommendation": {
      "description": "A recommendation is a list of recommended arbitrations to be emitted by the system\n",
      "properties": {
        "proposals": {
          "description": "The recommendation values mapping rule.name -\u003e formatted proposal",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Proposal"
          }
        },
        "timestamp": {
          "description": "The recommendation timestamp",
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "Rule": {
      "description": "An rule is a contraint expression that checked in the system against\n",
      "properties": {
        "expr": {
          "description": "The expression to match or query to make",
          "type": "string",
          "example": "100 * (px_volume_usage_bytes / px_volume_capacity_bytes) \u003e 80"
        },
        "for": {
          "description": "The duration/interval the expression must be valid for in seconds",
          "type": "integer",
          "format": "int64",
          "example": 3600
        },
        "issue": {
          "description": "The issue template",
          "type": "string",
          "example": "Portworx volume {{$labels.volume}} usage on {{$labels.host}} is high."
        },
        "name": {
          "description": "the rule description",
          "type": "string"
        },
        "proposal": {
          "description": "The proposal template",
          "type": "string",
          "example": "Add additional storage node to {{$labels.cluster}}"
        },
        "severity": {
          "type": "string",
          "enum": [
            "warning",
            "error",
            "critical"
          ]
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "BadRequest",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "invalid parameter"
        }
      }
    },
    "NotFound": {
      "description": "NotFound",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "object not found"
        }
      }
    },
    "ServerError": {
      "description": "ServerError",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "internal server error"
        }
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "examples": {
        "application/json": {
          "message": "access denied"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "security": [
    {
      "basicAuth": []
    }
  ]
}`))
}
