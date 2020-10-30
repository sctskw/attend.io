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
          "system"
        ],
        "responses": {
          "200": {
            "description": "api status",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/SystemStatus"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/attendee": {
      "get": {
        "tags": [
          "attendees"
        ],
        "operationId": "getAttendeeByField",
        "parameters": [
          {
            "type": "string",
            "format": "email",
            "description": "Attendee Email",
            "name": "email",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "Attendee ID",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "find an attendee by id",
            "schema": {
              "$ref": "#/definitions/Attendee"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/attendees/{eventId}": {
      "get": {
        "tags": [
          "attendees"
        ],
        "operationId": "getAttendeesByEventId",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Event ID",
            "name": "eventId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "list the attendees per event",
            "schema": {
              "$ref": "#/definitions/AttendeesList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events": {
      "get": {
        "tags": [
          "events"
        ],
        "responses": {
          "200": {
            "description": "list the events",
            "schema": {
              "$ref": "#/definitions/EventList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events/{id}": {
      "get": {
        "tags": [
          "events"
        ],
        "operationId": "getEventById",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Event ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the requested Event",
            "schema": {
              "$ref": "#/definitions/Event"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
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
              "$ref": "#/definitions/TalkList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Attendee": {
      "type": "object",
      "required": [
        "name_first",
        "name_last"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name_display": {
          "type": "string",
          "minLength": 1
        },
        "name_first": {
          "type": "string",
          "minLength": 1
        },
        "name_last": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "AttendeesList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Attendee"
      }
    },
    "Error": {
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
    "Event": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "talk": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "EventList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Event"
      }
    },
    "SystemStatus": {
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
    "Talk": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "attendees": {
          "$ref": "#/definitions/AttendeesList"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "TalkList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Talk"
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
          "system"
        ],
        "responses": {
          "200": {
            "description": "api status",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/SystemStatus"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/attendee": {
      "get": {
        "tags": [
          "attendees"
        ],
        "operationId": "getAttendeeByField",
        "parameters": [
          {
            "type": "string",
            "format": "email",
            "description": "Attendee Email",
            "name": "email",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "Attendee ID",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "find an attendee by id",
            "schema": {
              "$ref": "#/definitions/Attendee"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/attendees/{eventId}": {
      "get": {
        "tags": [
          "attendees"
        ],
        "operationId": "getAttendeesByEventId",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Event ID",
            "name": "eventId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "list the attendees per event",
            "schema": {
              "$ref": "#/definitions/AttendeesList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events": {
      "get": {
        "tags": [
          "events"
        ],
        "responses": {
          "200": {
            "description": "list the events",
            "schema": {
              "$ref": "#/definitions/EventList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events/{id}": {
      "get": {
        "tags": [
          "events"
        ],
        "operationId": "getEventById",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Event ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the requested Event",
            "schema": {
              "$ref": "#/definitions/Event"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
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
              "$ref": "#/definitions/TalkList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Attendee": {
      "type": "object",
      "required": [
        "name_first",
        "name_last"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name_display": {
          "type": "string",
          "minLength": 1
        },
        "name_first": {
          "type": "string",
          "minLength": 1
        },
        "name_last": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "AttendeesList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Attendee"
      }
    },
    "Error": {
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
    "Event": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "talk": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "EventList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Event"
      }
    },
    "SystemStatus": {
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
    "Talk": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "attendees": {
          "$ref": "#/definitions/AttendeesList"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "TalkList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Talk"
      }
    }
  }
}`))
}
