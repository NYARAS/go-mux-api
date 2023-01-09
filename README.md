# Go Mux API Demo
Simple Go demo code. Use Go Mux to create an API server

It's a simple API used to demo how to deploy a production grade Golang API on [AWS Elastic Container Service](https://aws.amazon.com/ecs/).

## Project prerequisites
- [Docker](https://docs.docker.com/install/)
- [docker-compose](https://docs.docker.com/compose/install/)

## Run locally
To run these commands you must be in the repository’s root.
### Environment variables

In order for you to run the application locally on your machine, you will need the following environment variables


### Environment Variables

  * `APP_DB_USERNAME` - Database username
  * `APP_DB_PASSWORD` - Database password
  * `APP_DB_NAME` - Database name

You can `export` these env or create a `env.sh` file and source it.

Example
```bash
export APP_DB_USERNAME=app
export APP_DB_PASSWORD=pleasechangeme
export APP_DB_NAME=api
```
### Docker-compose - Without Nginx Proxy

Build
```bash
docker-compose build
```
Run
```bash
docker-compose up
```
Server is listening on localhost:8080
### Docker-compose - With Nginx Proxy

Build
```bash
docker-compose -f docker-compose-proxy.yml build
```
Run
```bash
docker-compose -f docker-compose-proxy.yml up
```
Server is listening on localhost:8000

## License

Copyright (c) 2023 Calvine Otieno

Distributed under the MIT License. See the file LICENSE.
