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
    "version": "0.0.1"
  },
  "paths": {
    "/attendees": {
      "get": {
        "tags": [
          "attendees"
        ],
        "summary": "Retrieve an Attendee by Email or ID",
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
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "attendees"
        ],
        "summary": "Create an Attendee",
        "parameters": [
          {
            "description": "the Attendee to create",
            "name": "attendee",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Attendee"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Attendee was created",
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
    "/attendees/{id}": {
      "delete": {
        "tags": [
          "attendees"
        ],
        "summary": "Deletes an Attendee by ID",
        "operationId": "deleteAttendeeById",
        "parameters": [
          {
            "type": "string",
            "description": "Attendee ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Attendee was deleted."
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
    "/status": {
      "get": {
        "tags": [
          "system"
        ],
        "summary": "Retrieve the System Status",
        "responses": {
          "200": {
            "description": "API is operating normally",
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
    "/talks": {
      "get": {
        "tags": [
          "talks"
        ],
        "summary": "Retrieve All Talks",
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
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "talks"
        ],
        "summary": "Create a Talk",
        "parameters": [
          {
            "description": "the Talk to create",
            "name": "talk",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Talk"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Talk was created",
            "schema": {
              "$ref": "#/definitions/Talk"
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
    "/talks/{id}": {
      "get": {
        "tags": [
          "talks"
        ],
        "summary": "Retrieve a Talk by ID",
        "operationId": "getTalkById",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the requested talk",
            "schema": {
              "$ref": "#/definitions/Talk"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "talks"
        ],
        "summary": "Deletes a Talk by ID",
        "operationId": "deleteTalkById",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Talk was deleted."
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
    "/talks/{id}/attendees": {
      "get": {
        "tags": [
          "talks"
        ],
        "summary": "Retrieve the Attendees for a Talk",
        "operationId": "getTalkAttendees",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the list of attendees from the event",
            "schema": {
              "$ref": "#/definitions/AttendeeList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "patch": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "talks"
        ],
        "summary": "Add an Attendee to a Talk",
        "operationId": "addAttendeeToTalk",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "attendees",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Attendee was created",
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
    "/talks/{id}/attendees/{attendeeId}": {
      "delete": {
        "tags": [
          "talks"
        ],
        "summary": "Remove an Attendee from a Talk",
        "operationId": "deleteAttendeeFromTalk",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Attendee ID",
            "name": "attendeeId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Attendeed was deleted from Talk."
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
        },
        "ref": {
          "type": "string",
          "readOnly": true
        },
        "ref_talks": {
          "$ref": "#/definitions/TalkList"
        }
      }
    },
    "AttendeeList": {
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
        "ID": {
          "type": "string",
          "readOnly": true
        },
        "date_time_end": {
          "type": "string",
          "format": "date-time"
        },
        "date_time_start": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "presenter": {
          "type": "string",
          "minLength": 1
        },
        "ref_attendees": {
          "$ref": "#/definitions/AttendeeList"
        }
      }
    },
    "TalkList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Talk"
      }
    }
  },
  "x-google-allow": "all"
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
    "version": "0.0.1"
  },
  "paths": {
    "/attendees": {
      "get": {
        "tags": [
          "attendees"
        ],
        "summary": "Retrieve an Attendee by Email or ID",
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
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "attendees"
        ],
        "summary": "Create an Attendee",
        "parameters": [
          {
            "description": "the Attendee to create",
            "name": "attendee",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Attendee"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Attendee was created",
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
    "/attendees/{id}": {
      "delete": {
        "tags": [
          "attendees"
        ],
        "summary": "Deletes an Attendee by ID",
        "operationId": "deleteAttendeeById",
        "parameters": [
          {
            "type": "string",
            "description": "Attendee ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Attendee was deleted."
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
    "/status": {
      "get": {
        "tags": [
          "system"
        ],
        "summary": "Retrieve the System Status",
        "responses": {
          "200": {
            "description": "API is operating normally",
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
    "/talks": {
      "get": {
        "tags": [
          "talks"
        ],
        "summary": "Retrieve All Talks",
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
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "talks"
        ],
        "summary": "Create a Talk",
        "parameters": [
          {
            "description": "the Talk to create",
            "name": "talk",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Talk"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Talk was created",
            "schema": {
              "$ref": "#/definitions/Talk"
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
    "/talks/{id}": {
      "get": {
        "tags": [
          "talks"
        ],
        "summary": "Retrieve a Talk by ID",
        "operationId": "getTalkById",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the requested talk",
            "schema": {
              "$ref": "#/definitions/Talk"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "talks"
        ],
        "summary": "Deletes a Talk by ID",
        "operationId": "deleteTalkById",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Talk was deleted."
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
    "/talks/{id}/attendees": {
      "get": {
        "tags": [
          "talks"
        ],
        "summary": "Retrieve the Attendees for a Talk",
        "operationId": "getTalkAttendees",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the list of attendees from the event",
            "schema": {
              "$ref": "#/definitions/AttendeeList"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "patch": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "talks"
        ],
        "summary": "Add an Attendee to a Talk",
        "operationId": "addAttendeeToTalk",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "attendees",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Attendee was created",
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
    "/talks/{id}/attendees/{attendeeId}": {
      "delete": {
        "tags": [
          "talks"
        ],
        "summary": "Remove an Attendee from a Talk",
        "operationId": "deleteAttendeeFromTalk",
        "parameters": [
          {
            "type": "string",
            "description": "Talk ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Attendee ID",
            "name": "attendeeId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Attendeed was deleted from Talk."
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
        },
        "ref": {
          "type": "string",
          "readOnly": true
        },
        "ref_talks": {
          "$ref": "#/definitions/TalkList"
        }
      }
    },
    "AttendeeList": {
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
        "ID": {
          "type": "string",
          "readOnly": true
        },
        "date_time_end": {
          "type": "string",
          "format": "date-time"
        },
        "date_time_start": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "presenter": {
          "type": "string",
          "minLength": 1
        },
        "ref_attendees": {
          "$ref": "#/definitions/AttendeeList"
        }
      }
    },
    "TalkList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Talk"
      }
    }
  },
  "x-google-allow": "all"
}`))
}
