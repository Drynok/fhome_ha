# Futurehome backend software engineer home assignment.

Task definition:

For your home assignment, we want you to build a Go web server that
will perform the following tasks:
- accepts input data as JSON via an HTTP endpoint (/feed)
- shards the above data in such a way that the number of shards
does not exceed 5
- data from every shard is processed by exactly one worker from
a worker pool; the worker pool has the following
characteristics:
- any worker has a unique identifier
- no worker lives for more than 2 minutes (this means there
is a chance no worker is alive at some point)
- the initial number of workers is 3 and this number cannot
exceed 4
- processing the data means writing it to disk in batches
of 5 input items (in a file named using the worker
identifier)
- exposes an HTTP endpoint (/history) which returns a list of
worker identifiers and the number of processed messages for
each worker

### Folder structure
1. `/packages` - contains 
2. `/env` - contains configuration for environment variables.
3. `/handlers` - holds HTTP handlers.
5. `/client` - contains HTTP client and input for testing.
6. `/middleware` - contains middleware.

### Getting started:

1. Create `.env` file in the projet root with following content:

```bash
SERVER_MODE=release
SERVER_PORT=5050
```

1. To build the service run:
```bash
make build
```

1. To start the service run:
```bash
make up
```

1. To shutdown the service run:
```bash
make down
```

### Doing development:
1. Run unit tests:
```bash
go test ./... -v
```

1. You can find `Swagger API` documentation in `docs.yaml``.


1. Linting. Install `golangci-lint` [installed](https://golangci-lint.run/usage/install/). After that you can run:

   ```bash
   golangci-lint run
   ```
