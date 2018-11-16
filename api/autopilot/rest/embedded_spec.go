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
  "paths": {
    "/collectors": {
      "get": {
        "description": "Returns an array of telemetry collectors defined in the system",
        "tags": [
          "collector"
        ],
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
      },
      "post": {
        "description": "Create a new telemetry collector from the provided definition",
        "tags": [
          "collector"
        ],
        "summary": "Create a new telemetry collector",
        "operationId": "collectorCreate",
        "parameters": [
          {
            "description": "The collector to create",
            "name": "collector",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Collector"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Collector"
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
      }
    },
    "/collectors/{collector_id}": {
      "get": {
        "description": "Returns the request collected object",
        "tags": [
          "collector"
        ],
        "summary": "Get a list of telemetry collectors",
        "operationId": "collectorGet",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The id of the collector",
            "name": "collector_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Collector"
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
          "404": {
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
      "put": {
        "description": "Update the properties of the specified collector",
        "tags": [
          "collector"
        ],
        "summary": "Update a collector object",
        "operationId": "collectorUpdate",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The id of the collector",
            "name": "collector_id",
            "in": "path",
            "required": true
          },
          {
            "description": "The collector to create",
            "name": "collector",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Collector"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
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
      "delete": {
        "description": "Returns the request collected object",
        "tags": [
          "collector"
        ],
        "summary": "Get a list of telemetry collectors",
        "operationId": "collectorDelete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The id of the collector",
            "name": "collector_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          },
          "404": {
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
      "description": "A collector pulls data from a telemetry provider, parses, \nand reformats the data to be consumed by the autopilot engine.\n",
      "properties": {
        "id": {
          "description": "The collector id",
          "type": "string",
          "format": "uuid"
        },
        "params": {
          "description": "json data object",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "provider": {
          "description": "The telemetry provider for the collector",
          "type": "string",
          "enum": [
            "prometheus"
          ]
        }
      },
      "example": {
        "id": "046b6c7f-0b8a-43b9-b35d-6489e6daee91",
        "params": {
          "key": "{}"
        },
        "provider": "prometheus"
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
  "paths": {
    "/collectors": {
      "get": {
        "description": "Returns an array of telemetry collectors defined in the system",
        "tags": [
          "collector"
        ],
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
      },
      "post": {
        "description": "Create a new telemetry collector from the provided definition",
        "tags": [
          "collector"
        ],
        "summary": "Create a new telemetry collector",
        "operationId": "collectorCreate",
        "parameters": [
          {
            "description": "The collector to create",
            "name": "collector",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Collector"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Collector"
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
      }
    },
    "/collectors/{collector_id}": {
      "get": {
        "description": "Returns the request collected object",
        "tags": [
          "collector"
        ],
        "summary": "Get a list of telemetry collectors",
        "operationId": "collectorGet",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The id of the collector",
            "name": "collector_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Collector"
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
          "404": {
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
      "put": {
        "description": "Update the properties of the specified collector",
        "tags": [
          "collector"
        ],
        "summary": "Update a collector object",
        "operationId": "collectorUpdate",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The id of the collector",
            "name": "collector_id",
            "in": "path",
            "required": true
          },
          {
            "description": "The collector to create",
            "name": "collector",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Collector"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
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
      "delete": {
        "description": "Returns the request collected object",
        "tags": [
          "collector"
        ],
        "summary": "Get a list of telemetry collectors",
        "operationId": "collectorDelete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The id of the collector",
            "name": "collector_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          },
          "404": {
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
      "description": "A collector pulls data from a telemetry provider, parses, \nand reformats the data to be consumed by the autopilot engine.\n",
      "properties": {
        "id": {
          "description": "The collector id",
          "type": "string",
          "format": "uuid"
        },
        "params": {
          "description": "json data object",
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "provider": {
          "description": "The telemetry provider for the collector",
          "type": "string",
          "enum": [
            "prometheus"
          ]
        }
      },
      "example": {
        "id": "046b6c7f-0b8a-43b9-b35d-6489e6daee91",
        "params": {
          "key": "{}"
        },
        "provider": "prometheus"
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
