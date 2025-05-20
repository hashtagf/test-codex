# Hexagonal Demo

This repository contains a small demonstration of a Go project organised using
[hexagonal architecture](https://alistair.cockburn.us/hexagonal-architecture/).

The core business logic is defined in the `internal` package and accessed via
interfaces (ports). Adapters expose the behaviour to the outside world.

## Running the server

```
go run ./cmd/server
```

Then open `http://localhost:8080?name=Codex` in your browser to see the
response.
