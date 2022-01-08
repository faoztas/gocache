# GoCache

GoCache is an in-memory cache project. GoCache has an RWMutex and endless storage Job. GoCache is written with the go standard library.

### Prerequisites

* Go 1.17+

## Installation

You can clone this repo, then Run and Build.

```bash
$ cd $GOPATH
$ mkdir github.com
$ cd github.com
$ git clone https://github.com/faoztas/gocache.git
$ cd gocache                # env.json template is filled.
$ go mod tidy               # This command allows you to fetch all the dependencies that you need for testing in your module.
$ go build -o gocache       # Or $ go run .
$ ./gocache
```

## API Usage

### Get Key

Return JSON data from cache about value by key.

**Request:**

```
GET {{url}}/get?key=hello
```

**Response:** 200 OK

```json
{
  "data": "world!",
  "message": ""
}
```

**Response:** 400 Bad Request

```json
{
  "data": null,
  "message": "missing key"
}
```

**Response:** 404 Not Found

```json
{
  "data": null,
  "message": "record not found"
}
```

### Set Key

Save value by key. Return JSON data from cache about value by key.

**Request:**

```
POST {{url}}/set?key=hello
```

**Request Body:**

```
world!
```

**Response:** 200 OK

```json
{
  "data": "world!",
  "message": ""
}
```

**Response:** 400 Bad Request

```json
{
  "data": null,
  "message": "missing key"
}
```

### Delete Key

Return JSON data about delete status by key.

**Request:**

```
DELETE {{url}}/delete?key=hello
```

**Response:** 200 OK

```json
{
  "data": true,
  "message": ""
}
```

**Response:** 400 Bad Request

```json
{
  "data": null,
  "message": "missing key"
}
```

**Response:** 404 Not Found

```json
{
  "data": null,
  "message": "record not found"
}
```

### Flush Keys

Return JSON data about flush status.

**Request:**

```
DELETE {{url}}/flush
```

**Response:** 200 OK

```json
{
  "data": true,
  "message": ""
}
```

### General Errors

Using wrong method

**Request:**

```
GET {{url}}/flush
```

**Response:** 404 Not Found

```json
{
  "data": null,
  "message": "prefix /flush not found in GET method"
}
```

Using wrong path

**Request:**

```
GET {{url}}/
```

**Response:** 404 Not Found

```json
{
  "data": null,
  "message": "not found"
}
```

## Docker

**build:**

```bash
$ docker build . -t gocache
```

**run:**

```bash
$ docker run -dp 8000:8000 -i -t gocache
```

## Test

Run the tests

```bash
$ go test ./... -v
$ go test -cover ./... -v  # run tests with code coverage
```