## About

2FA PoC API is a proof-of-concept project that demonstrates a simple implementation of a Two-Factor Authentication (2FA) API. It provides endpoints for user registration, login, 2FA code generation and validation, and updating user credentials. The project is built using Go and utilizes the Chi router.

## Features

- User registration: Allows users to create an account by providing their credentials.
- User login: Allows users to authenticate using their registered credentials.
- Two-Factor Authentication (2FA): Generates and validates 2FA codes for additional security.
- User credential update: Enables users to update their account credentials.

## Prerequisites

Before running the project, ensure that the following dependencies are installed:

- Go (1.16 or later)
- Docker
- Docker Compose

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/alexandredsa/2fa-poc-api.git
   cd 2fa-poc-api
   ```

2. Set up the environment variables:
   
   Create a `.env` file in the project root directory and configure the following variables:

   ```shell
   POSTGRES_HOST=postgres
   POSTGRES_PORT=5432
   POSTGRES_USER=your_postgres_user
   POSTGRES_PASSWORD=your_postgres_password
   POSTGRES_DB=your_postgres_database
   REDIS_HOST=redis
   REDIS_PORT=6379
   ```

3. Build and start the Docker containers:

   ```shell
   docker-compose up --build
   ```

4. The API should now be running on `http://localhost:8080`. You can test the endpoints using tools like cURL, Insomnia, Postman, etc.


## Makefile Commands

The following commands are available in the Makefile:

- `build`: Build the application.
- `run`: Run the application.
- `test`: Run the tests.
- `test-coverage-html`: Generate test coverage HTML report.
- `docker-build`: Build the Docker image.
- `docker-run`: Run the Docker container.
- `docker-stop`: Stop and remove the Docker container.
- `docker-clean`: Remove the Docker container (if stopped).
- `docker-up`: Start the application using Docker Compose.
- `docker-down`: Stop and remove the Docker containers created by Docker Compose.

To execute a command, run `make <command>` in the terminal, where `<command>` is the desired command from the list above. For example, to build the application, run `make build`.


## API Endpoints

The following endpoints are available:

### Authentication

- `POST /auth/register` - Register a new user.
- `POST /auth/login` - Authenticate user credentials.
- `POST /auth/2fa/{component}/request` - Request a 2FA code for a specific component.
- `POST /auth/2fa/{component}/validate` - Validate a 2FA code for a specific component.
- `PUT /auth/credentials` - Update user credentials.
- `PUT /auth/{component}` - Update component data.
- `POST /auth/{component}/validate` - Validate component data.

### Account Data

- `PUT /account/credentials` - Update user credentials.

### 2FA Data

- `PUT /2fa/{component}` - Update component data.
- `POST /2fa/{component}/validate` - Validate component data.

