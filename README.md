This readme file provides an overview and documentation for the astrodata_api application's RESTful API, focusing on the package handlers.

The API provides access to various space weather data retrieval endpoints:

### Endpoints:

- ##### Base URL: https://i_don't_have_yet/api :(

- ##### Health check

```http
GET /health
```
- Description: checks the health of the API

- Response:

```http
status Code: 200 OK
content-Type: application/json
```

### Dst data

- Description: returns dst data for n period

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
GET /api/dst/by-date?date=data (YYYYMM format, Example: 202306)
```

```http
GET /dst/now/strength
```

- Response:
```http
Status Code: 200 OK
Content-Type: application/json
```

### Predictions (60% sure)

- Description: calculate and returns the predicted dst index for the next n hours (future predict (60% sure))

```http
GET /predict/6h
```

```http
GET /predict/1d
```

- Response:

```http
Status Code: 200 OK
Content-Type: application/json
```

### Bz data

- Description: retrieve bz data for n period

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

### Plasma temperature data

- Description: retrieve time plasma temperature data for n time period

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

- Error Responses

```http
Status Code: 400 Bad Request
Content-Type: application/json
```

- Example error Response:

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

### CORS Configs

- The API is configured to allow Cross-Origin Resource Sharing (CORS) by using the github.com/gin-contrib/cors middleware, enabling access to the API from different domains.

- The API will be available at https://somesite.xyz, make requests to the defined endpoints to access space weather data.

- Please ensure that the appropriate URL is used along with any required query parameters when making requests to the API endpoints.

## Contributing

- Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](./LICENSE)
