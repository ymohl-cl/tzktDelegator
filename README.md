# TZKT Delegator

TZKT Delegator is a project that provides services in Go for handling delegations on the Tezos blockchain network. The project consists of two main applications:

* API: A RESTful API that exposes a /xtz/delegations endpoint to provide delegations on the network. The delegations are limited to 100 items max per request. The API is implemented using the Echo framework and can be found in the cmd/api directory.
* Delegator Ingestor: A websocket client that subscribes to delegation events and ingests them into a Postgres database. If a reorg occurs, the data in the database is updated accordingly. The ingestor can be found in the cmd/delegator_injestor directory.
* History: A service to get the historical data from the TzKT API and store it in the database. The history can be found in the cmd/history directory. /!\ This service is helper to valid the fonctionnal test and do not implement best practice. Process is not elegant and pull + 700000 entries from TzKT API to record them in the database. :sweat_smile:

The project uses Postgres as the database and SQLC to generate the Go resources. A Makefile and a Docker Compose file are also included for ease of use.

## Getting Started

To get started with TZKT Delegator, follow these steps:

Clone the repository:

``` bash
<$ git clone https://github.com/ymohl-cl/tzktDelegator.git
```

Start the services:

``` bash
<$ docker compose up
```

test api status on `http://localhost:4242/healthcheck` and get delegations on `http://localhost:4242/xtz/delegations`

Run `make build` to build the project (all command). You can also build individual service with `make build.api`, `make build.history` and `make build.delegator_ingestor`.

## Suggestions for Improvement

Here are some suggestions for improving TZKT Delegator:

* Complete unit tests for all packages
* Add a service to ingest historical data with the TzKT API
* Add pagination to the xtz/delegations request
* /!\ Record the reorgs with timestamps to avoid to ingest the older data reorganized in a multi instance environment
* /!\ Check if reorgs exist for the current new data entry (reorg stategy in multi instance environment)
* Replace Echo with Gin (optional) for potential performance gains
* Add error handling for HTTP parameter parsing in the API
* Add Swagger documentation for the API
* Implement sql bulk record

## Folder Structure

The project has the following folder structure:

├── cmd
│   ├── api
│   │   └── main.go
│   └── delegator_injestor
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   └── logger..
├── internal
│   └── service..
├── README.md
├── sql
│   ├── migrations
│   │   └── ...
│   └── queries.sql
├── License
└── sqlc.yaml

## License

TZKT Delegator is licensed under the MIT License. See the LICENSE file for more information.

Contributions are welcome! If you have any suggestions or improvements, please open an issue or submit a pull request.
