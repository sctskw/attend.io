---
title: attend.io v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
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

```http
GET / HTTP/1.1

Accept: application/github.com/sctskw/attend.io.v1+json

```

```javascript

const headers = {
  'Accept':'application/github.com/sctskw/attend.io.v1+json'
};

fetch('/',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/github.com/sctskw/attend.io.v1+json'
}

result = RestClient.get '/',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/github.com/sctskw/attend.io.v1+json'
}

r = requests.get('/', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/github.com/sctskw/attend.io.v1+json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/github.com/sctskw/attend.io.v1+json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

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

```http
GET /talks HTTP/1.1

Accept: application/github.com/sctskw/attend.io.v1+json

```

```javascript

const headers = {
  'Accept':'application/github.com/sctskw/attend.io.v1+json'
};

fetch('/talks',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/github.com/sctskw/attend.io.v1+json'
}

result = RestClient.get '/talks',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/github.com/sctskw/attend.io.v1+json'
}

r = requests.get('/talks', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/github.com/sctskw/attend.io.v1+json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/talks', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/talks");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/github.com/sctskw/attend.io.v1+json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/talks", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /talks`

> Example responses

> 200 Response

```json
[
  {
    "id": 0,
    "name": "string"
  }
]
```

<h3 id="get__talks-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|list the talks|Inline|
|default|Default|generic error response|[Error](#schemaerror)|

<h3 id="get__talks-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Talk](#schematalk)]|false|none|none|
|» id|integer(int64)|false|read-only|none|
|» name|string|true|none|none|

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
  "id": 0,
  "name": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(int64)|false|read-only|none|
|name|string|true|none|none|

<h2 id="tocS_Attendee">Attendee</h2>
<!-- backwards compatibility -->
<a id="schemaattendee"></a>
<a id="schema_Attendee"></a>
<a id="tocSattendee"></a>
<a id="tocsattendee"></a>

```json
{
  "id": 0,
  "name": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(uuid)|false|read-only|none|
|name|string|true|none|none|

<h2 id="tocS_Event">Event</h2>
<!-- backwards compatibility -->
<a id="schemaevent"></a>
<a id="schema_Event"></a>
<a id="tocSevent"></a>
<a id="tocsevent"></a>

```json
{
  "id": 0,
  "name": "string",
  "talk": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(uuid)|false|read-only|none|
|name|string|true|none|none|
|talk|integer(uuid)|false|none|none|

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

