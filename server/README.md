# Todo Golang Server

## REST API

GET /todos -- list of todos -- 200, 404, 500
POST /todos -- create todo -- 204, 4xx
PATCH /todos/:id -- partially update todo by id -- 204/200, 404, 400, 500
DELETE /todos/:id -- delete todo by id -- 204, 404, 400

## Start:
1. `cd server`
2. `go run cmd/main/app.go`