# astro data api

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## structure:

```go
.
├── api
│   ├── handlers.go
│   ├── health
│   │   └── health.go
│   └── models.go
├── core
│   ├── core.go
│   └── parser.go
├── db
│   ├── db.go
│   └── models.go
├── go.mod
├── go.sum
├── main.go
├── README.md
└── vars
    └── vars.go
```

## api docs

- this readme file provides an overview and documentation for the astrodata_api application's RESTful API, 
  focusing on the package handlers

- the API provides access to various space weather data retrieval endpoints:

### endpoints:

- ##### base URL: https://i_dony_have_yet/api

- ##### health check

```http
GET /health
```
- description: checks the health of the API

- response:

```http
status Code: 200 OK
content-Type: application/json
```

### dst data

- description: returns dst data for n period

```http
GET /dst/7d
```

```http
GET /dst/current-month
```

```http
GET /dst/last-month
```

```http
GET /dst/now
```

```http
GET /api/dst/by-date?date=data (YYYYMM format, example: 202306)
```

```http
GET /dst/now/strength
```

- response:
```http
Status Code: 200 OK
Content-Type: application/json
```

### predictions (60% sure)

- description: calculate and returns the predicted dst index for the next n hours (future predict (60% sure))

```http
GET /predict/6h
```

```http
GET /predict/1d
```

- response:

```http
Status Code: 200 OK
Content-Type: application/json
```

### bz data

- description: retrieve bz data for n period

```http
GET /bz/6h
```

```http
GET /bz/1d
```

```http
GET /bz/3d
```

```http
GET /bz/7d
```

```http
GET /api/bz/7d
```

```http
GET /bz/now
```

```http
GET /bz/now/strength
```

### plasma temperature data

- description: retrieve time plasma temperature data for n time period

```http
GET /plasma/now
```

```http
GET /plasma/2h
```

```http
GET /plasma/6h
```

```http
GET /plasma/1d
```

```http
GET /api/plasma/1d
```

```http
GET /plasma/3d
```

```http
GET /api/plasma/3d
```

```http
GET /plasma/7d
```

- error responses

```http
Status Code: 400 Bad Request
Content-Type: application/json
```

- example error response:

```json
{
  "error": "date parameter is missing!"
}
```

```json
{
  "error": "Invalid date format"
}
```

### CORS configs

- the API is configured to allow Cross-Origin Resource Sharing (CORS) by using the github.com/gin-contrib/cors middleware, enabling access to the API from different domains

- the API will be available at https://somesite.xyz, make requests to the defined endpoints to access space weather data

- please ensure that the appropriate URL is used along with any required query parameters when making requests to the API endpoints

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

- please make sure to update tests as appropriate

## license

- [MIT](https://choosealicense.com/licenses/mit/)
