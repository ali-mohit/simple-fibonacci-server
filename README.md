# Simple Fibonacci Server (SFS)

Simple Fibonacci Server is a Go application that calculates Fibonacci numbers. It supports both command-line and HTTP server modes, providing a versatile way to interact with Fibonacci number calculations.

## Features

- **Command-line Interface**: Calculate Fibonacci numbers directly from the command line.
- **HTTP Server**: A RESTful API server that accepts requests to calculate Fibonacci numbers.
- **In-Memory Caching**: Improves performance by caching results of previous calculations.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.21.6 or higher
- Docker (optional for containerization)

### Installing

First, clone the repository to your local machine:

```bash
git clone https://github.com/ali-mohit/simple-fibonacci-server.git
cd simple-fibonacci-server
```
Then, build the application using the Makefile:
```bash
make fibo_app
```

## Running the Application in console
Command-line Interface

To calculate a Fibonacci number using the command line:
```bash
./fibo_app fibonacci [number]
```
Replace [number] with the Fibonacci sequence index you want to calculate.

## Running the Application on http server
To start the HTTP server:

```bash
./fibo_app serve
```

The server will start listening on http://localhost:8080. You can make requests to /fibonacci endpoint to calculate Fibonacci numbers.

## Using Docker
Build the Docker image:

```bash
docker build -t simple-fibonacci-server .
```

Run the Docker container:
```bash
docker-compose up
```

## API Reference
`POST:` `/fibonacci`

`Content-Type: application/json or application/x-msgpack`

Request Body:
```json
{
  "n": 10
}
```
Response:
```json
{
  "n": 55
}
```

## For testing
To run test:

```bash
make test
```

## License

This project is licensed under the MIT License - see the `LICENSE.md` file for details.