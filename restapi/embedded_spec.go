// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
    "application/github.com/sctskw/attend.io.v1+json"
  ],
  "produces": [
    "application/github.com/sctskw/attend.io.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple talk attendance API",
    "title": "attend.io",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "status"
        ],
        "responses": {
          "200": {
            "description": "api status",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/talks": {
      "get": {
        "tags": [
          "talks"
        ],
        "responses": {
          "200": {
            "description": "list the talks",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/talk"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "attendee": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "event": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "talk": {
          "type": "integer",
          "format": "uuid"
        }
      }
    },
    "status": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "talk": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/github.com/sctskw/attend.io.v1+json"
  ],
  "produces": [
    "application/github.com/sctskw/attend.io.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple talk attendance API",
    "title": "attend.io",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "status"
        ],
        "responses": {
          "200": {
            "description": "api status",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/talks": {
      "get": {
        "tags": [
          "talks"
        ],
        "responses": {
          "200": {
            "description": "list the talks",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/talk"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "attendee": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "event": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "talk": {
          "type": "integer",
          "format": "uuid"
        }
      }
    },
    "status": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "talk": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
