# FavouritesService

This service manages the watch history a user can have for a streaming app

## Characteristics
* It offers a full CRUD for every user's watch history.
* It creates a new history element for every movie or episode a user has watched.
* It uses the Chi router for a lightweight, yet easy to code API.

## Configuration
The app uses environment variables to build the database connection URL. If there was no environment variable, it will take a preset default.

```
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=default
DB_PASSWORD=example
DB_DATABASE=streamingdb
```

## Setup
To start your server:

* Run `go mod download` to download the dependencies
* Run `go build -ldflags="-s -w" server.exe` to compile the code into a static binary
* Start the endpoint by the generated `server.exe` binary file

Now the server will be active at [`localhost:7500`](http://localhost:7500).

