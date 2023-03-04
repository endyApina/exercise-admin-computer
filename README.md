# EXERCISE ADMIN COMPUTER

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=main)](https://travis-ci.org/joemccann/dillinger)

This repository provides an endpoint to keep track of computers issues by the company
This services run on port `9119` with the base HTTP URL `/api/computer` 

### Dependencies
- Install && Run Docker [docker](https://docs.docker.com/get-docker/)
- Install and run postgress with Docker
    ```sh
    docker pull postgres
     docker run -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=secret-password -d postgres
    ```