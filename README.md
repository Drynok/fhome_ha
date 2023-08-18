# Futurehome backend software engineer home assignment

This projet contains a draft of game of proximity back-end based on gRPC service. 

### Folder structure
1. `/packages` - contains
2. `/env` - environment variables
3. `/handlers` - HTTP handlers
4. `/libraries` - wrappers for third-party integartions like Redis

### Getting started:

1. Create `.env` file in the projet root with following content:

```bash
SERVER_MODE=release
SERVER_PORT=5050
```

1. To start the service run
```bash
make up
```

1. Open

### Doing development:
1. Run unit tests:
```bash
go test ./... -v
```