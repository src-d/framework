# Framework

## Description

Base package for src-d codebase. This package contain generic tools to be used
in different core packages for different parts of our platform.

So far, contains the following packages:

* **Configurable** standarizes the way to create containers.
* **Datbase** uses `configurable` to model database connection
  configurations.
* **Queue** defines interfaces for working with queue and implements them for
  AMQP and an in-memory queue.
