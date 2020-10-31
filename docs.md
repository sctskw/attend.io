---
title: attend.io v1.0.0
language_tabs:
  - shell: Shell
language_clients:
  - shell: ""
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="attend-io">attend.io v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

A simple talk attendance API

<h1 id="attend-io-system">system</h1>

## get__

> Code samples

```shell
# You can also use wget
curl -X GET / \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /`

*retrieve the system status*

> Example responses

> 200 Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="get__-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|api status|[SystemStatus](#schemasystemstatus)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="attend-io-talks">talks</h1>

## get__talks

> Code samples

```shell
# You can also use wget
curl -X GET /talks \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /talks`

*retrieve all Talks*

> Example responses

> 200 Response

```json
[
  {
    "ID": "string",
    "name": "string",
    "presenter": "string",
    "description": "string",
    "date_time_start": "2019-08-24T14:15:22Z",
    "date_time_end": "2019-08-24T14:15:22Z",
    "ref_attendees": [
      {
        "id": "string",
        "ref": "string",
        "name_display": "string",
        "name_first": "string",
        "name_last": "string",
        "email": "user@example.com"
      }
    ]
  }
]
```

<h3 id="get__talks-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|list the talks|[TalkList](#schematalklist)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## post__talks

> Code samples

```shell
# You can also use wget
curl -X POST /talks \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`POST /talks`

*create a Talk*

> Body parameter

```json
{
  "name": "string",
  "presenter": "string",
  "description": "string",
  "date_time_start": "2019-08-24T14:15:22Z",
  "date_time_end": "2019-08-24T14:15:22Z",
  "ref_attendees": [
    {
      "name_display": "string",
      "name_first": "string",
      "name_last": "string",
      "email": "user@example.com"
    }
  ]
}
```

<h3 id="post__talks-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Talk](#schematalk)|false|the Talk to create|

> Example responses

> 200 Response

```json
{
  "ID": "string",
  "name": "string",
  "presenter": "string",
  "description": "string",
  "date_time_start": "2019-08-24T14:15:22Z",
  "date_time_end": "2019-08-24T14:15:22Z",
  "ref_attendees": [
    {
      "id": "string",
      "ref": "string",
      "name_display": "string",
      "name_first": "string",
      "name_last": "string",
      "email": "user@example.com"
    }
  ]
}
```

<h3 id="post__talks-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Talk was created|[Talk](#schematalk)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getTalkById

<a id="opIdgetTalkById"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /talks/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /talks/{id}`

*retrieve a Talk by ID*

<h3 id="gettalkbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Talk ID|

> Example responses

> 200 Response

```json
{
  "ID": "string",
  "name": "string",
  "presenter": "string",
  "description": "string",
  "date_time_start": "2019-08-24T14:15:22Z",
  "date_time_end": "2019-08-24T14:15:22Z",
  "ref_attendees": [
    {
      "id": "string",
      "ref": "string",
      "name_display": "string",
      "name_first": "string",
      "name_last": "string",
      "email": "user@example.com"
    }
  ]
}
```

<h3 id="gettalkbyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|the requested talk|[Talk](#schematalk)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## deleteTalkById

<a id="opIddeleteTalkById"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /talks/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`DELETE /talks/{id}`

*Deletes a Talk by ID*

<h3 id="deletetalkbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Talk ID|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="deletetalkbyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|Talk was deleted.|None|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getTalkAttendees

<a id="opIdgetTalkAttendees"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /talks/{id}/attendees \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /talks/{id}/attendees`

*Retrieve the Attendees for a Talk*

<h3 id="gettalkattendees-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Talk ID|

> Example responses

> 200 Response

```json
[
  {
    "id": "string",
    "ref": "string",
    "name_display": "string",
    "name_first": "string",
    "name_last": "string",
    "email": "user@example.com"
  }
]
```

<h3 id="gettalkattendees-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|the list of attendees from the event|[AttendeeList](#schemaattendeelist)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## addAttendeeToTalk

<a id="opIdaddAttendeeToTalk"></a>

> Code samples

```shell
# You can also use wget
curl -X PATCH /talks/{id}/attendees \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`PATCH /talks/{id}/attendees`

*add an Attendee to a Talk*

> Body parameter

```json
[
  "string"
]
```

<h3 id="addattendeetotalk-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Talk ID|
|body|body|array[string]|false|none|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "ref": "string",
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}
```

<h3 id="addattendeetotalk-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Attendee was created|[Attendee](#schemaattendee)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## deleteAttendeeFromTalk

<a id="opIddeleteAttendeeFromTalk"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /talks/{id}/attendees/{attendeeId} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`DELETE /talks/{id}/attendees/{attendeeId}`

*Deletes a Talk with specific ID*

<h3 id="deleteattendeefromtalk-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Talk ID|
|attendeeId|path|string|true|Attendee ID|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="deleteattendeefromtalk-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|Attendeed was deleted from Talk.|None|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="attend-io-attendees">attendees</h1>

## getAttendeeByField

<a id="opIdgetAttendeeByField"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /attendees \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /attendees`

*retrieve an Attendee by its fields*

<h3 id="getattendeebyfield-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|email|query|string(email)|false|Attendee Email|
|id|query|string|false|Attendee ID|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "ref": "string",
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}
```

<h3 id="getattendeebyfield-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|find an attendee by id|[Attendee](#schemaattendee)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## post__attendees

> Code samples

```shell
# You can also use wget
curl -X POST /attendees \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`POST /attendees`

*create an Attendee*

> Body parameter

```json
{
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}
```

<h3 id="post__attendees-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Attendee](#schemaattendee)|false|the Attendee to create|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "ref": "string",
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}
```

<h3 id="post__attendees-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Attendee was created|[Attendee](#schemaattendee)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## deleteAttendeeById

<a id="opIddeleteAttendeeById"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /attendees/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`DELETE /attendees/{id}`

*Deletes an Attendee by ID*

<h3 id="deleteattendeebyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Attendee ID|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="deleteattendeebyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|Attendee was deleted.|None|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Talk">Talk</h2>
<!-- backwards compatibility -->
<a id="schematalk"></a>
<a id="schema_Talk"></a>
<a id="tocStalk"></a>
<a id="tocstalk"></a>

```json
{
  "ID": "string",
  "name": "string",
  "presenter": "string",
  "description": "string",
  "date_time_start": "2019-08-24T14:15:22Z",
  "date_time_end": "2019-08-24T14:15:22Z",
  "ref_attendees": [
    {
      "id": "string",
      "ref": "string",
      "name_display": "string",
      "name_first": "string",
      "name_last": "string",
      "email": "user@example.com"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|ID|string|false|read-only|none|
|name|string|true|none|none|
|presenter|string|false|none|none|
|description|string|false|none|none|
|date_time_start|string(date-time)|false|none|none|
|date_time_end|string(date-time)|false|none|none|
|ref_attendees|[AttendeeList](#schemaattendeelist)|false|none|none|

<h2 id="tocS_TalkList">TalkList</h2>
<!-- backwards compatibility -->
<a id="schematalklist"></a>
<a id="schema_TalkList"></a>
<a id="tocStalklist"></a>
<a id="tocstalklist"></a>

```json
[
  {
    "ID": "string",
    "name": "string",
    "presenter": "string",
    "description": "string",
    "date_time_start": "2019-08-24T14:15:22Z",
    "date_time_end": "2019-08-24T14:15:22Z",
    "ref_attendees": [
      {
        "id": "string",
        "ref": "string",
        "name_display": "string",
        "name_first": "string",
        "name_last": "string",
        "email": "user@example.com"
      }
    ]
  }
]

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Talk](#schematalk)]|false|none|none|

<h2 id="tocS_Attendee">Attendee</h2>
<!-- backwards compatibility -->
<a id="schemaattendee"></a>
<a id="schema_Attendee"></a>
<a id="tocSattendee"></a>
<a id="tocsattendee"></a>

```json
{
  "id": "string",
  "ref": "string",
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|read-only|none|
|ref|string|false|read-only|none|
|name_display|string|false|none|none|
|name_first|string|true|none|none|
|name_last|string|true|none|none|
|email|string(email)|false|none|none|

<h2 id="tocS_AttendeeList">AttendeeList</h2>
<!-- backwards compatibility -->
<a id="schemaattendeelist"></a>
<a id="schema_AttendeeList"></a>
<a id="tocSattendeelist"></a>
<a id="tocsattendeelist"></a>

```json
[
  {
    "id": "string",
    "ref": "string",
    "name_display": "string",
    "name_first": "string",
    "name_last": "string",
    "email": "user@example.com"
  }
]

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Attendee](#schemaattendee)]|false|none|none|

<h2 id="tocS_Error">Error</h2>
<!-- backwards compatibility -->
<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "code": 0,
  "message": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int64)|false|none|none|
|message|string|true|none|none|

<h2 id="tocS_SystemStatus">SystemStatus</h2>
<!-- backwards compatibility -->
<a id="schemasystemstatus"></a>
<a id="schema_SystemStatus"></a>
<a id="tocSsystemstatus"></a>
<a id="tocssystemstatus"></a>

```json
{
  "code": 0,
  "message": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int64)|false|none|none|
|message|string|true|none|none|

