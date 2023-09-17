# Prometheus and Grafana Observability Experiment

This project is designed to provide hands-on experience with Prometheus and Grafana for observability within a simplified API environment.

## Overview

The project demonstrates how different Prometheus metric types can be employed within an API without the complexities of web frameworks or libraries beyond the core Prometheus components. It includes an instrumented API, a Prometheus server, and a Grafana server, all running in a local development environment.

## Prerequisites

Before you begin, ensure you have the following prerequisites:

- Docker
- Docker Compose

## Getting Started

To set up the development environment, follow these steps:

1. Run `make dev-up` to start the environment.
2. Access Prometheus server on http://localhost:9090
3. Access Grafana server on http://localhost:3000
4. Log in to the local Grafana server using username and password as `admin`.
5. Add a new Prometheus datasource to Grafana. This set up works on Mac using http://prometheus:9090 as the server URL.
6. Access the development API on http://localhost:8000 to start recording custom metrics. See [server.go](./pkg/server.go) for defined endpoints.

## Stopping the Environment

To stop the development environment, run `make dev-down`.
