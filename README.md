# EXERCISE ADMIN COMPUTER

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This repository provides an endpoint to keep track of computers issues by the company
This services run on port `9119` with the base HTTP URL `/api/computer` 

### Dependencies
- Install && Run Docker [docker](https://docs.docker.com/get-docker/)
- Install and run postgress with Docker
    ```sh
    docker pull postgres
    docker run -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=secret-password -d postgres
    ```

### Environment Varables
```go
service_name: "computer-records"
service_port: "9119"
database_host: "localhost"
database_port: "5432"
database_name: "greenbone"
database_password: "PASSWORD"
```

NOTE: `CREATE POSTGRES DATABASE`

To build the service run

```sh
go build -o admin-computer
```

Import postman collection

```sh
download && import Greenbone.postman_collection.json
```

Afterwards the service can be started with `./admin-computer`.