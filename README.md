# datacollector

Push data via UDP (port 8080).
Pull pushed data via `GET /metrics` (port 9090)

This is just a proof of concept atm!

## Build

```shell
go mod download # get the dependencies
go build -o bin/datacollector
```

## Example

### Run datacollector

```shell
bin/datacollector
```

### Send stuff via UDP

```shell
nc -u 127.0.0.1 8080
```

### Get the stuff

```shell
curl 127.0.0.1:9090/metrics
```
