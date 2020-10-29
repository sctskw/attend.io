# attend.io
A simple talk attendance API

## Version: 1.0.0

### /

#### GET
##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | api status | [SystemStatus](#systemstatus) |
| default | generic error response | [Error](#error) |

### /talks

#### GET
##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | list the talks | [ [Talk](#talk) ] |
| default | generic error response | [Error](#error) |

### Models

#### Talk

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | long |  | No |
| name | string |  | Yes |

#### Attendee

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer (uuid) |  | No |
| name | string |  | Yes |

#### Event

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer (uuid) |  | No |
| name | string |  | Yes |
| talk | integer (uuid) |  | No |

#### Error

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | long |  | No |
| message | string |  | Yes |

#### SystemStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | long |  | No |
| message | string |  | Yes |
