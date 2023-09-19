# Example Loans Service

## Overview
This repository contains an example service showcasing various components commonly found in a Go application. It is designed to serve as an illustrative example for building a service, covering areas such as gRPC APIs, data validation, database interactions, prometheus metrics, opentracing implementation, and more.

## Features
gRPC API: Demonstrates the implementation of gRPC APIs for communication.
Data Validation: Shows how to validate data using the go-ozzo/ozzo-validation library.
Database Interaction: Illustrates database interactions with the GORM library.
Interceptors: Includes examples of gRPC interceptors for various purposes.
Error Handling: Demonstrates error handling and reporting strategies.
Configuration: Highlights how to manage configuration using environment variables.
Getting Started
Follow these instructions to set up and run the example service locally.

## Prerequisites
[Go](https://go.dev/doc/install)

[Mysql Setup](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/macos-installation-pkg.html)

[Jaeger](https://www.jaegertracing.io/docs/1.49/getting-started/#:~:text=It%20includes%20the%20Jaeger%20UI,(a%20single%20command%20line).&text=You%20can%20then%20navigate%20to,to%20access%20the%20Jaeger%20UI) - Setup would work without jaeger integration as well.

## Good

## Installation
### Clone the repository:

```
git clone https://github.com/yourusername/exampleservice.git
```

### Navigate to the project directory:

```
cd exampleservice
```
### Initial Setup:

#### Install go dependencies
```
go mod download
```

#### Install proto dependencies
```
make proto-deps
```
This will install all gRPC, proto, and buf dependencies

#### Create Database
Create the database that you will be using for this service and run the queries added In file `queries.sql` to create required tables

#### Setup Local configuration
create `dev.toml` by duplicating `default.toml`. And set the `APP_ENV` env variable as `dev` 
```
cp config/default.toml config/dev.toml
export APP_ENV="dev"
```
This toml file will be used to configure your local setup. separate dev.toml is created by every developer because configurations can differ in their local setups.
Update db configuration, jaeger configuration according to your local setup.

## Build and run the service:
1. Refresh RPC files
```
make proto-refresh
```
This command will delete existing rpc files and generate new ones according to the protos.

2. Build and Run the service
```
make go-build-api && ./bin/api
```
This command will build the service binary and run the service

## Testing

### Dependencies
```
make mock-deps
```
This command will generate required mocks for the UTs to run

### running tests
```
make test-unit
```
This command will run unit tests and show the output in terminal

## Prerequisites
A basic understanding of the following technologies can be beneficial for understanding and working with this code:

- [RPC](https://www.techtarget.com/searchapparchitecture/definition/Remote-Procedure-Call-RPC)
- [Protocol Buffers (Proto)](https://protobuf.dev/)
- [gRPC](https://betterprogramming.pub/understanding-grpc-60737b23e79e)
- [Prometheus](https://prometheus.io/docs/introduction/overview/)
- [OpenTracing](https://www.sentinelone.com/blog/what-is-opentracing/)
- [Jaeger](https://aws.amazon.com/what-is/jaeger/#:~:text=AWS%20App%20Mesh%3F-,What%20is%20Jaeger%3F,complete%20a%20single%20software%20function.)