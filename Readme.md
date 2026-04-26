# EventLens

EventLens is a distributed log platform demo that generates logs from mock services, routes them through Kafka, indexes them in Elasticsearch, and exposes them in a React dashboard for live monitoring and historical search.

## What's included

This project supports two main use cases:

- Live log streaming to the frontend through WebSocket.
- Historical log search through HTTP -> gRPC -> Elasticsearch.

The mock services currently covered are:

- auth-service
- order-service
- payment-service

## Architecture

The current architecture flow is:

- auth-service, order-service, and payment-service publish raw logs to Kafka (raw logs).
- log-processor consumes raw Kafka messages.
- log-processor indexes logs into Elasticsearch (indexes).
- After successful indexing, log-processor publishes the same event to Kafka (logs-live).
- api-gateway consumes Kafka (logs-live) and pushes live events to the frontend over WebSocket /ws.
- The frontend calls api-gateway over HTTP for search operations.
- api-gateway calls query-service over gRPC.
- query-service reads from Elasticsearch and returns filtered search results.

## Data Flow

### Live Log Streaming

- Mock services publish logs to Kafka (raw logs).
- log-processor consumes raw log topics.
- log-processor indexes logs into Elasticsearch.
- After successful indexing, log-processor publishes the same payload to Kafka (logs-live).
- api-gateway consumes Kafka (logs-live).
- api-gateway pushes those logs to WebSocket clients at /ws.
- The frontend renders them in the Live Logs tab.

### Search

- The frontend sends a POST request to api-gateway.
- api-gateway converts the JSON body into a gRPC request.
- api-gateway calls query-service.
- query-service builds Elasticsearch filters and executes the query.
- The response returns logs and cursor metadata in base_response.
- The frontend uses sorted_value for cursor-based pagination.

## Tech Stack

### Backend

- Go
- Fiber
- gRPC
- Kafka with Sarama
- Elasticsearch
- Gorilla WebSocket

#### Infra

- Docker Compose
- Kafka
- Zookeeper
- Elasticsearch
- Kibana
- Grafana

## Run all of the Services in parallel

- `make -j3`
