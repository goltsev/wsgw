# Websocket API Gateway

This application is a Websocket gateway for BitMex Websocket API.

## Configuration

Application uses `.env` file and environment variables for configuration.

List of used environment variables:

* `KEY` -- BitMex API key
* `SECRET` -- BitMex API secret
* `URL` -- BitMex Websocket API endpoint URL
* `EXPIRES` -- date of expiration for connections

## Launch

Application can be launched using:

```shell
go run .
```

There is also a very simple client for testing the application locally that can be run with:

```shell
go run ./client/
```
