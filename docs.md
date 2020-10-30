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

> Example responses

> 200 Response

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "name": "string",
    "attendees": [
      {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
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

## getTalkById

<a id="opIdgetTalkById"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /talks/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /talks/{id}`

<h3 id="gettalkbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string(uuid)|true|Talk ID|

> Example responses

> 200 Response

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "name": "string",
  "attendees": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
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

<h1 id="attend-io-events">events</h1>

## get__events

> Code samples

```shell
# You can also use wget
curl -X GET /events \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /events`

> Example responses

> 200 Response

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "name": "string",
    "talk": "0156e0a5-f7a8-4710-8864-e3389ea9b565"
  }
]
```

<h3 id="get__events-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|list the events|[EventList](#schemaeventlist)|
|default|Default|generic error response|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getEventById

<a id="opIdgetEventById"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /events/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /events/{id}`

<h3 id="geteventbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string(uuid)|true|Event ID|

> Example responses

> 200 Response

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "name": "string",
  "talk": "0156e0a5-f7a8-4710-8864-e3389ea9b565"
}
```

<h3 id="geteventbyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|the requested Event|[Event](#schemaevent)|
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
curl -X GET /attendee \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /attendee`

<h3 id="getattendeebyfield-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|email|query|string(email)|false|Attendee Email|
|id|query|string(uuid)|false|Attendee ID|

> Example responses

> 200 Response

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
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

## getAttendeesByEventId

<a id="opIdgetAttendeesByEventId"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /attendees/{eventId} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

`GET /attendees/{eventId}`

<h3 id="getattendeesbyeventid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|eventId|path|string(uuid)|true|Event ID|

> Example responses

> 200 Response

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "name_display": "string",
    "name_first": "string",
    "name_last": "string",
    "email": "user@example.com"
  }
]
```

<h3 id="getattendeesbyeventid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|list the attendees per event|[AttendeesList](#schemaattendeeslist)|
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "name": "string",
  "attendees": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
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
|id|string(uuid)|false|read-only|none|
|name|string|true|none|none|
|attendees|[AttendeesList](#schemaattendeeslist)|false|none|none|

<h2 id="tocS_TalkList">TalkList</h2>
<!-- backwards compatibility -->
<a id="schematalklist"></a>
<a id="schema_TalkList"></a>
<a id="tocStalklist"></a>
<a id="tocstalklist"></a>

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "name": "string",
    "attendees": [
      {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string(uuid)|false|read-only|none|
|name_display|string|false|none|none|
|name_first|string|true|none|none|
|name_last|string|true|none|none|
|email|string(email)|false|none|none|

<h2 id="tocS_AttendeesList">AttendeesList</h2>
<!-- backwards compatibility -->
<a id="schemaattendeeslist"></a>
<a id="schema_AttendeesList"></a>
<a id="tocSattendeeslist"></a>
<a id="tocsattendeeslist"></a>

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
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

<h2 id="tocS_Event">Event</h2>
<!-- backwards compatibility -->
<a id="schemaevent"></a>
<a id="schema_Event"></a>
<a id="tocSevent"></a>
<a id="tocsevent"></a>

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "name": "string",
  "talk": "0156e0a5-f7a8-4710-8864-e3389ea9b565"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string(uuid)|false|read-only|none|
|name|string|true|none|none|
|talk|string(uuid)|false|none|none|

<h2 id="tocS_EventList">EventList</h2>
<!-- backwards compatibility -->
<a id="schemaeventlist"></a>
<a id="schema_EventList"></a>
<a id="tocSeventlist"></a>
<a id="tocseventlist"></a>

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "name": "string",
    "talk": "0156e0a5-f7a8-4710-8864-e3389ea9b565"
  }
]

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Event](#schemaevent)]|false|none|none|

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

