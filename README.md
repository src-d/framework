# framework [![Build Status](https://travis-ci.org/src-d/framework.svg?branch=master)](https://travis-ci.org/src-d/framework) [![codecov.io](https://codecov.io/gh/src-d/framework/branch/master/graph/badge.svg?token=am2H6bJkdp)](https://codecov.io/gh/src-d/framework)

**framework** provides abstractions to services used across multiple projects.

## Services

* **configurable** standarizes the way to create configuration containers.
* **database** package provides access to SQL databases.
* **queue** provides access to message brokers.

## Development

Run tests with:

  go test -v ./...

Tests require the following services running:

* PostgreSQL
* RabbitMQ

They also need the `etcd` binary present in `PATH`.
