# RESTful API for ScratchPay Clinic Search

NOTE: The base Golang RESTful API backbone is forked from http://github.com/olguncengiz/scratchpay-clinicsearch

This project is designed to allow search in multiple clinic providers and display results from all the available clinics by any of the following:

* Clinic Name
* State [ex: "CA" or "California"]
* Availability [ex: from:09:00, to:20:00]

This is including search by multiple criteria in the same time like search by state and availability together.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.13 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

After installing Go and Docker, run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/olguncengiz/scratchpay-clinicsearch.git

cd scratchpay-clinicsearch

# run the RESTful API server
make run
```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `GET /healthcheck`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `POST /v1/login`: authenticates a user and generates a JWT
* `POST /v1/dentalClinics`: returns a list of the dental clinics with given search criteria
* `GET /v1/vetClinics`: returns a list of the vet clinics with given search criteria

Try the URL `http://localhost:8080/healthcheck` in a browser, and you should see something like `"OK v1.0.0"` displayed.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following 
more complex scenarios:

```shell
# authenticate the user via: POST /v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "pass"}' http://localhost:8080/v1/login
# should return a JWT token like: {"token":"...JWT token here..."}

# with the above JWT token, access the album resources, such as: POST /v1/dentalClinics
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ...JWT token here..." -d '{"state": "California"}' http://localhost:8080/v1/dentalClinics
# should return a list of dental clinic records in the JSON format
```

### Assumptions
Following is a list of assumptions made for this project.
* Before any search request, the user must authenticate using username and password (By default, the values are `"username": "demo", "password": "pass"`.
* The dental clinic search and vet clinic search will be made on different endpoints due to the different data resource files provided. 
* The **JSON** files for dental and vet clinic data must be in the root folder of project, where `server` file can be found.

## Project Layout

The project uses the following layout:
 
```
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── config               configuration files for different environments
├── internal             private application and library code
│   ├── auth             authentication feature
│   ├── clinics          clinics related library code
│   ├── config           configuration library
│   ├── entity           entity definitions and domain logic
│   ├── errors           error types and handling
│   ├── healthcheck      healthcheck feature
│   └── test             helpers for testing purpose
├── pkg                  public library code
│   ├── accesslog        access log middleware
│   ├── log              structured and context-aware logger
```

The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example, 
the `clinics` directory contains the application logic related with the clinic search feature. 

Within each feature package, code are organized in layers (API, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

### Managing Configurations

The application configuration is represented in `internal/config/config.go`. When the application starts,
it loads the configuration from a configuration file as well as environment variables. The path to the configuration 
file is specified via the `-config` command line argument which defaults to `./config/local.yml`. Configurations
specified in environment variables should be named with the `APP_` prefix and in upper case. When a configuration
is specified in both a configuration file and an environment variable, the latter takes precedence. 

The `config` directory contains the configuration files named after different environments. For example,
`config/local.yml` corresponds to the local development environment and is used when running the application 
via `make run`.

Do not keep secrets in the configuration files. Provide them via environment variables instead. Secrets can be populated from a secret
storage (e.g. HashiCorp Vault) into environment variables in a bootstrap script (e.g. `cmd/server/entryscript.sh`). 

## Deployment

The application can be run as a docker container. You can use `make build-docker` to build the application 
into a docker image. The docker container starts with the `cmd/server/entryscript.sh` script which reads 
the `APP_ENV` environment variable to determine which configuration file to use. For example,
if `APP_ENV` is `qa`, the application will be started with the `config/qa.yml` configuration file.

You can also run `make build` to build an executable binary named `server`. Then start the API server using the following
command,

```shell
./server -config=./config/prod.yml
```