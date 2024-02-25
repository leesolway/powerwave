# Powerwave

Powerwave is a tool for managing power meters. It provides both a command-line interface (CLI) and an HTTP server for interacting with power meter data.

## Features

- Get meters by customer name
- Get meter readings by serial ID and date
- Supports both CLI and HTTP interfaces

## Installation

To install Powerwave, you can either build it from source or use Docker.
In order to make this OS agnostic i would recommend building with docker.

### Docker

1. Clone the repository:

```bash
git clone https://github.com/leesolway/powerwave.git
cd powerwave
```

2. Build the docker image

```bash
docker build -t powerwave .
```

### Window ###

Although all the examples use Docker the application can be built without Docker in Windows providing the go environment has been setup.

```bash
GOOS=windows GOARCH=amd64 go build -o powerwave.exe ./src
```

### Linux ###

```bash
GOOS=linux GOARCH=amd64 go build -o powerwave ./src
```

## Usage

### CLI
Powerwave provides a CLI interface for managing power meters. Here are some examples of CLI commands:

#### Get Meters by Customer

```bash
# Get meters by customer name
docker run --rm powerwave /app/app getmeters <customer>
```

**Example**
```bash
docker run --rm powerwave /app/app getmeters "Albers Facilities Management"
```

***

#### Get meter readings by serial and date

```bash
# Get meter reading by serial ID and date
docker run --rm powerwave /app/app getreading <serialID> <date>
```

**Example**
 ```bash
 docker run --rm powerwave /app/app getreading 1111-1111-1111 2023-01-01
 ```

### HTTP Server
Powerwave includes an HTTP server. You can run the server using the following command:

```bash
docker run --rm -p 8080:8080 -e PORT=8080 powerwave
```

#### Get Meters by Customer

**Example Request**
```http
GET /meters/Albers%20Facilities%20Management
```

**Example Response**
```json
[
    {
        "SerialID": "1111-1111-3333",
        "Building": "Student Halls",
        "Customer": "Albers Facilities Management",
        "DailyKWh": 40
    }
]
```

#### Get meter readings by serial and date

**Example Request**
```http
GET /readings/1111-1111-1111/2023-01-01
```

**Example Response**
```json
{
    "reading": {
        "SerialID": "1111-1111-1111",
        "Date": "2023-01-01",
        "KWhForDay": 20,
        "KWhForMonth": 620
    },
    "serialID": "1111-1111-1111"
}
```

## Configuration

Powerwave can be configured using environment variables.

- **PORT**: The port number for the HTTP server (default: 8080)
- **DEBUG_PORT**: The port number for debugging (default: 8081)

## TODO

- Increase loggging, tracing and error handling cases
- Expand on the placeholder middleware
- GIN production mode
- Setup development enviroment including dev container, debugger, linting etc
- Add further tests
- Add tests to a build pipeline
- Validation and sanitization needs expanding
- API documentation such as OAS
- Add support for graceful shutdown
- Process monitoring in Docker
- Decouple the data source for mocking and testing