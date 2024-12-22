Calculation webservice

>It can evaluate mathematical expressions with *, /, +, -, () operators

On default, it is run on localhost:8080 (It can be changed manually if needed)

# Table of Contents

1. [API Endpoints](#api-endpoints)
   - [Request Schema](#request-schema)
   - [Allowed Methods](#allowed-methods)
   - [Responses](#responses)
2. [Launch Command](#launch-command)
3. [Example of Use](#example-of-use)
    - [Successful Request](#successful-request)
    - [Invalid Expression](#invalid-expression)
    - [Bad Request](#bad-request)
    - [Method Not Allowed](#method-not-allowed)
    - [Internal Server Error](#internal-server-error)
4. [Middleware](#middleware)
   - [Logger](#logger)

# API Endpoints
>/api/v1/calculation

App will simply calculate an expression and will give the result with precision of 2 decimal places (like 1.00)
## Request Schema
```http request
POST /api/v1/calculate
Content-Type: application/json

{
"expression": "expression"
}
```
## Allowed methods
> POST

## Responses
> 200
```json
{
  "result": "result"
}
```
> 400
```json
{
   "error": "Bad request"
}
```
> 422
```json
{
  "error": "Expression is not valid"
}
```
> 500
```json
{
  "error": "Internal server Error"
}
```

# Launch Command
```shell
go run ./cmd/main.go
```

# Example of Use
## Successful Request
### Curl Command
```shell
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
### Status Code
> 200
### Response
```json
{
  "result": "6"
}
```
## Bad Request
### Curl Command
```shell
curl --location 'localhost:8080/api/v1/calculate' \
-X POST \
--header 'Content-Type: application/json'
```
### Status Code
> 400
### Response
```json
{
  "error": "Bad request"
}
```

## Method not Allowed
### Curl Command
```shell
curl --location 'localhost:8080/api/v1/calculate'
```
### Status Code
> 405
### Response
```json
{
  "error": "Method not Allowed"
}
```
## Invalid Expression
### Curl Command
```shell
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*b"
}'
```
### Status Code
> 422
### Response
```json
{
  "error": "Expression is not valid"
}
```
## Internal Server Error
### Explanation
Due to the fact that internal server error occur is unpredictable, we are unable to provide an example
### Status Code
> 500
```json
{
   "error": "Internal Server Error"
}
```

# Middleware
## Logger
Logging middleware is used to log requests<br/>
It logs only the method