# Mock Coupon API

This repository contains a mock api that we may use to experiment.

## Prerequisites

1. GNU Make
2. Docker

## Setting up

1. Clone the repository.

2. Run the server using `make docker-compose/up`. Ensure docker engine is 
running before running this command.

```bash
make docker-compose/up
```

## API 

The API should be up and running in `http://localhost:8080` as soon as you
run the above commands.

- Send a `GET` request to fetch all the available offers.
- You can add an optional parameter `epoch` in the request data to get
the data only after the given epoch.

## Configuration

Check the `docker/docker-compose.yaml` to see all the available configurations
you can change.

