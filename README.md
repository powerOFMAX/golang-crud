# [Go/Docker] Gin-gonic & MongoDB Messages Backend

This is a basic CRUD sample using Gin-gonic and MongoDB. The project has a docker configuration if you don't want to install go globally.

For this project I used [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) But you can setup a local Mongo Database if you want to.

```
# messages collection

_id: ObjectId
title: String
content: String
author: String
```

## Tree
```
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── app
│       └── server.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── config
│   │   └── db.go
│   ├── controllers
│   │   └── messages.go
│   └── routes
│       └── routes.go
└── scripts
    ├── development.sh
    └── production.sh
```

## Available Endpoints
```
[GET]    /api/v1/ - Welcome (To test the Endpoint)
[GET]    /api/v1/messages - Get All Messages)
[POST]   /api/v1/message - (Create Message)
[GET]    /api/v1/message/:messageId - (Get Single Message)
[PUT]    /api/v1message/:messageId - (Edit Message)
[DELETE] /api/v1/message/:messageId - (Delete Message)
```
You can test the endpoints using Postman. I left a collection in the project.
# Setup the project
First you must setup your database credentials. You can use `.env-example` as reference.

```console
$ cp .env-example .env
```

## Using with Docker
Start production enviroment
```console
$ docker-compose up -d production
// View logs
$ docker-compose logs --tail 100 -f production
```

Start development enviroment 
```console
$ docker-compose up development
```

Re-building docker
```console
$ docker-compose build --no-cache
```

Attach to bash
```console
$ docker-compose exec <production|development> sh
```

## Example
After running docker-compose open:

- development: http://localhost:8080/

- production:  http://localhost:8081/ 

## Without docker
```bash
$ make serve
``` 
You may need to execute `go mod download` in `src` folder first

## Configure scripts
#### ./scripts/production.sh
```bash
cd src 
go mod download
go build -o /bin/app && /bin/app
```

#### ./scripts/development.sh`
```bash
cd src
go run main.go
```
