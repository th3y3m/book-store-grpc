# Welcome to Bookstore Microservices Repository

This repository contains a Bookstore microservices application that uses gRPC for service communication, Docker for containerization, Protobuf for data serialization, and an API Gateway to expose RESTful endpoints for external clients.

## Overview

### Microservices:
- **Book Service** (book_service): Manages book-related operations, such as listing books.
- **User Service** (user_service): Manages user-related operations, such as listing users. It also interacts with the Book Service to fetch related book data.
- **API Gateway** (api_gateway): A REST API gateway that translates incoming HTTP requests into gRPC calls to the microservices.

### Technologies:
- **gRPC**: A high-performance RPC framework used for communication between the microservices.
- **Protocol Buffers** (Protobuf): Used for defining service interfaces and message types in a language-neutral and platform-neutral way.
- **Docker**: Containerization platform used to deploy the microservices.
- **API Gateway**: A REST API facade for external clients, exposing REST endpoints to interact with gRPC services.

## Prerequisites

Before running the application, ensure you have the following installed:

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/your-username/bookstore-microservices.git
cd bookstore-microservices
```

### Build and Run the Application
To build and run the services, run the following command:

```bash
docker-compose up --build
```

This will:

- Spin up the postgres_service to handle database operations.
- Launch the Book Service, User Service, and the API Gateway.

#### Connect me via: truongtanhuy3006@gmail.com

##### &#169; 2024 th3y3m

