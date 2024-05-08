# Snappfood Test
## Running The Project
The simplest way to run this project is using docker compose(use docker-compose if you're on older version.)
    
    docker compose up

## Testing Routes:
Import postman.json to postman.

For info on what each route does Read [ROUTES.md](./ROUTES.md) file.

## Running Tests
There are 2 sample tests in this project.
- Unit Test (ordering/services/agents/agent_service_test.go)
- Intergation test (cmd/web/main_test.go)

**Note:** For running intergation test you will need to copy .env file to cmd/web folder.

**Note VSCODE:** If you are using vscode open ordering folder.

## Running locally with .env file:
- Make sure you have postgresql running somewhere.
- Make sure you have rabbitmq running somewhere.
- Copy `sample.env` as `.env` in same directory and update values.

Run this command to get tables created:

    go run cmd/migrate/main.go

At last run this command:

    go run cmd/web/main.go


## Parameters Checking and fault tolerance:
- Code in this repo is not doing extra work for checking system flow (Ex: Wont check If trip status was previously updated)