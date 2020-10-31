## attend.io

A Talk Attendance API

version: `0.0.1`

---

#####Looking for API Docs? Please visit the [Live Documentation](https://attend-io-294107.uc.r.appspot.com/docs)


Development
---

#### Requirements
 - GoLang >= 1.14.9
 - Swagger CodeGen: https://github.com/swagger-api/swagger-codegen
 - GCP Cloud Credential (please ask)
 
 
#### Start Local Server

```
make serve
```

*WARNING: you cannot run the server without GCP credentials locally*
```
panic: couldn't connect to firestore db: google: could not find default credentials. 
See https://developers.google.com/accounts/docs/application-default-credentials for more information.

```


#### Run Tests

```
make test

>
Running tests in ./services
PASS tests.TestExampleTestSuite/TestAttendeesAPI/create_an_Attendee (0.72s)
PASS tests.TestExampleTestSuite/TestAttendeesAPI/get_Attendee_by_ID (0.05s)
PASS tests.TestExampleTestSuite/TestAttendeesAPI/get_Attendee_by_Email (0.06s)
PASS tests.TestExampleTestSuite/TestAttendeesAPI/join_a_Talk (0.64s)
PASS tests.TestExampleTestSuite/TestAttendeesAPI/delete_an_Attendee (0.43s)
PASS tests.TestExampleTestSuite/TestAttendeesAPI (1.91s)
PASS tests.TestExampleTestSuite/TestSystemAPI_GetStatus (0.00s)
PASS tests.TestExampleTestSuite/TestTalksAPI/create_a_Talk (0.16s)
PASS tests.TestExampleTestSuite/TestTalksAPI/fetch_all_Talks (0.11s)
PASS tests.TestExampleTestSuite/TestTalksAPI/fetch_Talk_by_ID (0.04s)
PASS tests.TestExampleTestSuite/TestTalksAPI/add_Attendee_to_Talk (0.55s)
PASS tests.TestExampleTestSuite/TestTalksAPI/remove_an_Attendee_from_a_Talk (0.20s)
PASS tests.TestExampleTestSuite/TestTalksAPI/delete_a_Talk (0.21s)
PASS tests.TestExampleTestSuite/TestTalksAPI (1.27s)
PASS tests.TestExampleTestSuite (3.21s)
PASS tests (cached)
EMPTY services
```

#### Generate API Docs
```
make docs
```

#### View Local API Docs
```
PORT=8001 make server
```
Docs will be available here: http://127.0.0.1:8001/



Examples
---

`GET /status` 

*Retrieve the System Status*

```shell
curl -X GET https://attend-io-294107.uc.r.appspot.com/status \
    -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*
```json
{
  "code": 200,
  "message": "attend.io is alive"
}
```

---

`GET /talks` 

*Retrieve All Talks*

```shell
curl -X GET https://attend-io-294107.uc.r.appspot.com/talks \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*
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
---


`POST /talks`

*Create a Talk*

```shell
curl -X POST https://attend-io-294107.uc.r.appspot.com/talks \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Request Body*
```json
{
  "name": "string",
  "presenter": "string",
  "description": "string",
  "date_time_start": "2019-08-24T14:15:22Z",
  "date_time_end": "2019-08-24T14:15:22Z"
}
```

*Example Response*
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
---

`GET /talks/{id}`

*Retrieve a Talk by ID*

```shell
curl -X GET https://attend-io-294107.uc.r.appspot.com/talks/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*
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
---

`DELETE /talks/{id}`

*Deletes a Talk by ID*

```shell
curl -X DELETE https://attend-io-294107.uc.r.appspot.com/talks/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*
```json
{
  "code": 204,
  "message": "the Talk was deleted."
}
```

---

`GET /talks/{id}/attendees`

*Retrieve the Attendees for a Talk*
```shell
curl -X GET https://attend-io-294107.uc.r.appspot.com/talks/{id}/attendees \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*

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
---

`PATCH /talks/{id}/attendees`

*Add an Attendee to a Talk*

```shell
curl -X PATCH https://attend-io-294107.uc.r.appspot.com/talks/{id}/attendees \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Request Body*

```json
[
  "ID-1",
  "ID-2"
]
```

*Example Response*

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
---

`DELETE /talks/{id}/attendees/{attendeeId}`

*Remove an Attendee from a Talk*

```shell
curl -X DELETE https://attend-io-294107.uc.r.appspot.com/talks/{id}/attendees/{attendeeId} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*
```json
{
  "code": 204,
  "message": "the Attendee was removed from the Talk."
}
```
---

`GET /attendees`

*Retrieve an Attendee by Email or ID*

```shell
# You can also use wget
curl -X GET https://attend-io-294107.uc.r.appspot.com/attendees \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*
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
---

`POST /attendees`

*Create an Attendee*

```shell
curl -X POST https://attend-io-294107.uc.r.appspot.com/attendees \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Request Body*

```json
{
  "name_display": "string",
  "name_first": "string",
  "name_last": "string",
  "email": "user@example.com"
}
```

*Example Response*

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
---

`DELETE /attendees/{id}`

*Deletes an Attendee by ID*

```shell
curl -X DELETE https://attend-io-294107.uc.r.appspot.com/attendees/{id} \
  -H 'Accept: application/github.com/sctskw/attend.io.v1+json'

```

*Example Response*

```json
{
  "code": 204,
  "message": "the Attendee was deleted."
}
```
