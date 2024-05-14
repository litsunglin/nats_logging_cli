
This Go code uses the NATS Go client library to connect to a NATS server, subscribe to the "debug_logging" channel, and print received messages.

## How to build
```
git clone https://github.com/litsunglin/nats_logging_cli.git
cd nats_logging_cli
go get github.com/nats-io/nats.go
CGO_ENABLED=0 go build .
```
